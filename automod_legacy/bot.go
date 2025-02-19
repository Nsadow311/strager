package automod_legacy

import (
	"time"

	"github.com/Nsadow311/stranger/analytics"
	"github.com/Nsadow311/stranger/bot"
	"github.com/Nsadow311/stranger/bot/eventsystem"
	"github.com/Nsadow311/stranger/commands"
	"github.com/Nsadow311/stranger/common"
	"github.com/Nsadow311/stranger/common/pubsub"
	"github.com/Nsadow311/stranger/moderation"
	"github.com/jonas747/discordgo/v2"
	"github.com/jonas747/dstate/v4"
	"github.com/karlseguin/ccache"
)

var _ bot.BotInitHandler = (*Plugin)(nil)

var (
	// cache configs because they are used often
	confCache *ccache.Cache
)

func (p *Plugin) BotInit() {
	commands.MessageFilterFuncs = append(commands.MessageFilterFuncs, CommandsMessageFilterFunc)

	eventsystem.AddHandlerAsyncLastLegacy(p, HandleMessageUpdate, eventsystem.EventMessageUpdate)

	pubsub.AddHandler("update_automod_legacy_rules", HandleUpdateAutomodRules, nil)
	confCache = ccache.New(ccache.Configure().MaxSize(1000))
}

// Invalidate the cache when the rules have changed
func HandleUpdateAutomodRules(event *pubsub.Event) {
	confCache.Delete(KeyConfig(event.TargetGuildInt))
}

// CachedGetConfig either retrieves from local application cache or redis
func CachedGetConfig(gID int64) (*Config, error) {
	confItem, err := confCache.Fetch(KeyConfig(gID), time.Minute*5, func() (interface{}, error) {
		c, err := GetConfig(gID)
		if err != nil {
			return nil, err
		}

		// Compile sites and words
		c.Sites.GetCompiled()
		c.Words.GetCompiled()

		return c, nil
	})

	if err != nil {
		return nil, err
	}

	return confItem.Value().(*Config), nil
}

func CommandsMessageFilterFunc(evt *eventsystem.EventData, msg *discordgo.Message) bool {
	return !CheckMessage(evt, msg)
}

func HandleMessageUpdate(evt *eventsystem.EventData) {
	CheckMessage(evt, evt.MessageUpdate().Message)
}

func CheckMessage(evt *eventsystem.EventData, m *discordgo.Message) bool {
	if !bot.IsNormalUserMessage(m) {
		return false
	}

	if m.Author.ID == common.BotUser.ID || m.Author.Bot || m.GuildID == 0 {
		return false // Pls no panicerinos or banerinos self
	}

	if !evt.HasFeatureFlag(featureFlagEnabled) {
		return false
	}

	cs := evt.GS.GetChannelOrThread(m.ChannelID)
	if cs == nil {
		logger.WithField("channel", m.ChannelID).Error("Channel not found in state")
		return false
	}

	config, err := CachedGetConfig(cs.GuildID)
	if err != nil {
		logger.WithError(err).Error("Failed retrieving config")
		return false
	}

	if !config.Enabled {
		return false
	}

	member := dstate.MemberStateFromMember(m.Member)
	member.GuildID = m.GuildID

	del := false // Set if a rule triggered a message delete
	punishMsg := ""
	highestPunish := PunishNone
	muteDuration := 0

	rules := []Rule{config.Spam, config.Invite, config.Mention, config.Links, config.Words, config.Sites}

	didCheck := false

	// We gonna need to have this locked while we check
	for _, r := range rules {
		if r.ShouldIgnore(cs, m, member) {
			continue
		}
		didCheck = true
		d, punishment, msg, err := r.Check(m, cs)
		if d {
			del = true
		}
		if err != nil {
			logger.WithError(err).WithField("guild", cs.GuildID).Error("Failed checking aumod rule:", err)
			continue
		}

		// If the rule did not trigger a deletion there wasn't any violation
		if !d {
			continue
		}

		punishMsg += msg + "\n"

		if punishment > highestPunish {
			highestPunish = punishment
			muteDuration = r.GetMuteDuration()
		}
	}

	if !del {
		if didCheck {
			go analytics.RecordActiveUnit(cs.GuildID, &Plugin{}, "checked")
		}
		return false
	}

	go analytics.RecordActiveUnit(cs.GuildID, &Plugin{}, "rule_triggered")

	if punishMsg != "" {
		// Strip last newline
		punishMsg = punishMsg[:len(punishMsg)-1]
	}

	go func() {
		switch highestPunish {
		case PunishNone:
			err = moderation.WarnUser(nil, cs.GuildID, cs, m, common.BotUser, &member.User, "Automoderator: "+punishMsg)
		case PunishMute:
			err = moderation.MuteUnmuteUser(nil, true, cs.GuildID, cs, m, common.BotUser, "Automoderator: "+punishMsg, member, muteDuration)
		case PunishKick:
			err = moderation.KickUser(nil, cs.GuildID, cs, m, common.BotUser, "Automoderator: "+punishMsg, &member.User)
		case PunishBan:
			err = moderation.BanUser(nil, cs.GuildID, cs, m, common.BotUser, "Automoderator: "+punishMsg, &member.User)
		}

		// Execute the punishment before removing the message to make sure it's included in logs
		common.BotSession.ChannelMessageDelete(m.ChannelID, m.ID)

		if err != nil && err != moderation.ErrNoMuteRole && !common.IsDiscordErr(err, discordgo.ErrCodeMissingPermissions, discordgo.ErrCodeMissingAccess) {
			logger.WithError(err).Error("Error carrying out punishment")
		}
	}()

	return true

}
