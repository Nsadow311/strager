package commands

import (
	"fmt"
	"strings"

	"github.com/Nsadow311/stranger/bot/paginatedmessages"
	"github.com/Nsadow311/stranger/common"
	"github.com/jonas747/dcmd/v4"
	"github.com/jonas747/discordgo/v2"
)

var cmdHelp = &YAGCommand{
	Name:        "Help",
	Aliases:     []string{"commands", "h", "how", "command"},
	Description: "Shows help about all or one specific command",
	CmdCategory: CategoryGeneral,
	RunInDM:     true,
	Arguments: []*dcmd.ArgDef{
		{Name: "command", Type: dcmd.String},
	},

	RunFunc:  cmdFuncHelp,
	Cooldown: 10,
}

func CmdNotFound(search string) string {
	return fmt.Sprintf("Couldn't find command '%s'", search)
}

func cmdFuncHelp(data *dcmd.Data) (interface{}, error) {
	target := data.Args[0].Str()

	// Send the targetted help in the channel it was requested in
	resp := dcmd.GenerateTargettedHelp(target, data, data.ContainerChain[0], &dcmd.StdHelpFormatter{})
	for _, v := range resp {
		ensureEmbedLimits(v)
	}

	if target != "" {
		if len(resp) != 1 {
			// Send command not found in same channel
			return CmdNotFound(target), nil
		}

		// Send short help in same channel
		return resp, nil
	}

	// Send full help in DM
	ir, err := createInteractiveHelp(data.Author.ID, resp)
	if ir != nil || err != nil {
		return ir, err
	}

	if data.Source == dcmd.TriggerSourceDM {
		return nil, nil
	}

	return "You've got mail!", nil
}

func createInteractiveHelp(userID int64, helpEmbeds []*discordgo.MessageEmbed) (interface{}, error) {
	channel, err := common.BotSession.UserChannelCreate(userID)
	if err != nil {
		return "Something went wrong, maybe you have DMs disabled? I don't want to spam this channel so here's a external link to available commands: <https://docs.dastranger.xyz/commands>", err
	}

	// prepend a introductionairy first page
	firstPage := &discordgo.MessageEmbed{
		Title: "Stranger Help!",
		Description: `Stranger is a fully functional multi-purpose discord bot that is configured through the web interface at https://dastranger.xyz.
For more in depth help and information you should visit https://docs.dastranger.xyz/ as this command only shows information about commands.
		
		
**Use the emojis under to change pages**`,
	}

	var pageLayout strings.Builder
	for i, v := range helpEmbeds {
		pageLayout.WriteString(fmt.Sprintf("**Page %d**: %s\n", i+2, v.Title))
	}
	firstPage.Fields = []*discordgo.MessageEmbedField{
		{Name: "Help pages", Value: pageLayout.String()},
	}

	helpEmbeds = append([]*discordgo.MessageEmbed{firstPage}, helpEmbeds...)

	_, err = paginatedmessages.CreatePaginatedMessage(0, channel.ID, 1, len(helpEmbeds), func(p *paginatedmessages.PaginatedMessage, page int) (*discordgo.MessageEmbed, error) {
		embed := helpEmbeds[page-1]
		return embed, nil
	})
	if err != nil {
		return "Something went wrong, make sure you don't have the bot blocked!", err

	}

	return nil, nil
}
