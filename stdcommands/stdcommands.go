package stdcommands

import (
	"github.com/Nsadow311/stranger/bot"
	"github.com/Nsadow311/stranger/bot/eventsystem"
	"github.com/Nsadow311/stranger/commands"
	"github.com/Nsadow311/stranger/common"
	"github.com/Nsadow311/stranger/stdcommands/advice"
	"github.com/Nsadow311/stranger/stdcommands/allocstat"
	"github.com/Nsadow311/stranger/stdcommands/banserver"
	"github.com/Nsadow311/stranger/stdcommands/calc"
	"github.com/Nsadow311/stranger/stdcommands/catfact"
	"github.com/Nsadow311/stranger/stdcommands/ccreqs"
	"github.com/Nsadow311/stranger/stdcommands/createinvite"
	"github.com/Nsadow311/stranger/stdcommands/currentshard"
	"github.com/Nsadow311/stranger/stdcommands/currenttime"
	"github.com/Nsadow311/stranger/stdcommands/customembed"
	"github.com/Nsadow311/stranger/stdcommands/dcallvoice"
	"github.com/Nsadow311/stranger/stdcommands/define"
	"github.com/Nsadow311/stranger/stdcommands/dogfact"
	"github.com/Nsadow311/stranger/stdcommands/findserver"
	"github.com/Nsadow311/stranger/stdcommands/globalrl"
	"github.com/Nsadow311/stranger/stdcommands/guildunavailable"
	"github.com/Nsadow311/stranger/stdcommands/howlongtobeat"
	"github.com/Nsadow311/stranger/stdcommands/info"
	"github.com/Nsadow311/stranger/stdcommands/invite"
	"github.com/Nsadow311/stranger/stdcommands/leaveserver"
	"github.com/Nsadow311/stranger/stdcommands/listflags"
	"github.com/Nsadow311/stranger/stdcommands/listroles"
	"github.com/Nsadow311/stranger/stdcommands/memstats"
	"github.com/Nsadow311/stranger/stdcommands/ping"
	"github.com/Nsadow311/stranger/stdcommands/poll"
	"github.com/Nsadow311/stranger/stdcommands/roll"
	"github.com/Nsadow311/stranger/stdcommands/setstatus"
	"github.com/Nsadow311/stranger/stdcommands/simpleembed"
	"github.com/Nsadow311/stranger/stdcommands/sleep"
	"github.com/Nsadow311/stranger/stdcommands/statedbg"
	"github.com/Nsadow311/stranger/stdcommands/stateinfo"
	"github.com/Nsadow311/stranger/stdcommands/throw"
	"github.com/Nsadow311/stranger/stdcommands/toggledbg"
	"github.com/Nsadow311/stranger/stdcommands/topcommands"
	"github.com/Nsadow311/stranger/stdcommands/topevents"
	"github.com/Nsadow311/stranger/stdcommands/topgames"
	"github.com/Nsadow311/stranger/stdcommands/topic"
	"github.com/Nsadow311/stranger/stdcommands/topservers"
	"github.com/Nsadow311/stranger/stdcommands/unbanserver"
	"github.com/Nsadow311/stranger/stdcommands/undelete"
	"github.com/Nsadow311/stranger/stdcommands/viewperms"
	"github.com/Nsadow311/stranger/stdcommands/weather"
	"github.com/Nsadow311/stranger/stdcommands/wouldyourather"
	"github.com/Nsadow311/stranger/stdcommands/xkcd"
	"github.com/Nsadow311/stranger/stdcommands/yagstatus"
)

var (
	_ bot.BotInitHandler       = (*Plugin)(nil)
	_ commands.CommandProvider = (*Plugin)(nil)
)

type Plugin struct{}

func (p *Plugin) PluginInfo() *common.PluginInfo {
	return &common.PluginInfo{
		Name:     "Standard Commands",
		SysName:  "standard_commands",
		Category: common.PluginCategoryCore,
	}
}

func (p *Plugin) AddCommands() {
	commands.AddRootCommands(p,
		// Info
		info.Command,
		invite.Command,

		// Standard
		define.Command,
		weather.Command,
		calc.Command,
		topic.Command,
		catfact.Command,
		dogfact.Command,
		advice.Command,
		ping.Command,
		throw.Command,
		roll.Command,
		customembed.Command,
		simpleembed.Command,
		currenttime.Command,
		listroles.Command,
		memstats.Command,
		wouldyourather.Command,
		poll.Command,
		undelete.Command,
		viewperms.Command,
		topgames.Command,
		xkcd.Command,
		howlongtobeat.Command,

		// Maintenance
		stateinfo.Command,
		leaveserver.Command,
		banserver.Command,
		allocstat.Command,
		unbanserver.Command,
		topservers.Command,
		topcommands.Command,
		topevents.Command,
		currentshard.Command,
		guildunavailable.Command,
		yagstatus.Command,
		setstatus.Command,
		createinvite.Command,
		findserver.Command,
		dcallvoice.Command,
		ccreqs.Command,
		sleep.Command,
		toggledbg.Command,
		globalrl.Command,
		listflags.Command,
	)

	statedbg.Commands()

}

func (p *Plugin) BotInit() {
	eventsystem.AddHandlerAsyncLastLegacy(p, ping.HandleMessageCreate, eventsystem.EventMessageCreate)
}

func RegisterPlugin() {
	common.RegisterPlugin(&Plugin{})
}
