package info

import (
	"fmt"

	"github.com/Nsadow311/stranger/commands"
	"github.com/Nsadow311/stranger/common"
	"github.com/jonas747/dcmd/v4"
)

var Command = &commands.YAGCommand{
	CmdCategory: commands.CategoryGeneral,
	Name:        "Info",
	Description: "Responds with bot information",
	RunInDM:     true,
	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		info := fmt.Sprintf(`**Stranger - A Fully Functional Multi-Purpose Discord Bot**
This bot focuses on being configurable and therefore is one of the most advanced bots.
It can perform a range of general purpose functionality (Reddit feeds, various commands, moderation utilities, automoderator functionality and so on) and it's configured through a web control panel.
The bot is run by Nsado311 but is open source (<https://github.com/Nsadow311/stranger>), so if you know Go and want to make some contributions, feel free to make a PR. (This is a fork of YAGPDB)
Control panel: <https://%s/manage>
				`, common.ConfHost.GetString())

		return info, nil
	},
}
