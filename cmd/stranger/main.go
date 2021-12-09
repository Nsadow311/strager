package main

import (
	"github.com/Nsadow311/stranger/analytics"
	"github.com/Nsadow311/stranger/common/featureflags"
	"github.com/Nsadow311/stranger/common/prom"
	"github.com/Nsadow311/stranger/common/run"
	"github.com/Nsadow311/stranger/web/discorddata"

	// Core yagpdb packages

	"github.com/Nsadow311/stranger/admin"
	"github.com/Nsadow311/stranger/bot/paginatedmessages"
	"github.com/Nsadow311/stranger/common/internalapi"
	"github.com/Nsadow311/stranger/common/scheduledevents2"

	// Plugin imports
	"github.com/Nsadow311/stranger/automod"
	"github.com/Nsadow311/stranger/automod_legacy"
	"github.com/Nsadow311/stranger/autorole"
	"github.com/Nsadow311/stranger/aylien"
	"github.com/Nsadow311/stranger/cah"
	"github.com/Nsadow311/stranger/commands"
	"github.com/Nsadow311/stranger/customcommands"
	"github.com/Nsadow311/stranger/discordlogger"
	"github.com/Nsadow311/stranger/logs"
	"github.com/Nsadow311/stranger/moderation"
	"github.com/Nsadow311/stranger/notifications"
	"github.com/Nsadow311/stranger/premium"
	"github.com/Nsadow311/stranger/premium/patreonpremiumsource"
	"github.com/Nsadow311/stranger/reddit"
	"github.com/Nsadow311/stranger/reminders"
	"github.com/Nsadow311/stranger/reputation"
	"github.com/Nsadow311/stranger/rolecommands"
	"github.com/Nsadow311/stranger/rsvp"
	"github.com/Nsadow311/stranger/safebrowsing"
	"github.com/Nsadow311/stranger/serverstats"
	"github.com/Nsadow311/stranger/soundboard"
	"github.com/Nsadow311/stranger/stdcommands"
	"github.com/Nsadow311/stranger/streaming"
	"github.com/Nsadow311/stranger/tickets"
	"github.com/Nsadow311/stranger/timezonecompanion"
	"github.com/Nsadow311/stranger/twitter"
	"github.com/Nsadow311/stranger/verification"
	"github.com/Nsadow311/stranger/youtube"
	"github.com/Nsadow311/stranger/twitch"
	// External plugins
	
)

func main() {

	run.Init()

	//BotSession.LogLevel = discordgo.LogInformational
	paginatedmessages.RegisterPlugin()
	discorddata.RegisterPlugin()

	// Setup plugins
	analytics.RegisterPlugin()
	safebrowsing.RegisterPlugin()
	discordlogger.Register()
	commands.RegisterPlugin()
	stdcommands.RegisterPlugin()
	serverstats.RegisterPlugin()
	notifications.RegisterPlugin()
	customcommands.RegisterPlugin()
	reddit.RegisterPlugin()
	moderation.RegisterPlugin()
	reputation.RegisterPlugin()
	aylien.RegisterPlugin()
	streaming.RegisterPlugin()
	automod_legacy.RegisterPlugin()
	automod.RegisterPlugin()
	logs.RegisterPlugin()
	autorole.RegisterPlugin()
	reminders.RegisterPlugin()
	soundboard.RegisterPlugin()
	youtube.RegisterPlugin()
	rolecommands.RegisterPlugin()
	cah.RegisterPlugin()
	tickets.RegisterPlugin()
	verification.RegisterPlugin()
	premium.RegisterPlugin()
	patreonpremiumsource.RegisterPlugin()
	scheduledevents2.RegisterPlugin()
	twitter.RegisterPlugin()
	rsvp.RegisterPlugin()
	timezonecompanion.RegisterPlugin()
	admin.RegisterPlugin()
	internalapi.RegisterPlugin()
	prom.RegisterPlugin()
	featureflags.RegisterPlugin()
	twitch.RegisterPlugin()
	run.Run()
}
