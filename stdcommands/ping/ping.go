package ping

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Nsadow311/stranger/bot/eventsystem"
	"github.com/Nsadow311/stranger/commands"
	"github.com/Nsadow311/stranger/common"
	"github.com/jonas747/dcmd/v4"
)

var Command = &commands.YAGCommand{
	CmdCategory:     commands.CategoryDebug,
	Name:            "Ping",
	Description:     "Shows the latency from the bot to the discord servers.",
	LongDescription: "Note that high latencies can be the fault of ratelimits and the bot itself, it's not a absolute metric.",

	RunFunc: func(data *dcmd.Data) (interface{}, error) {
		return fmt.Sprintf(":PONG;%d", time.Now().UnixNano()), nil
	},
}

func HandleMessageCreate(evt *eventsystem.EventData) {
	m := evt.MessageCreate()

	bUser := common.BotUser
	if bUser == nil {
		return
	}

	if bUser.ID != m.Author.ID {
		return
	}

	// ping pong
	split := strings.Split(m.Content, ";")
	if split[0] != ":PONG" || len(split) < 2 {
		return
	}

	parsed, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return
	}

	taken := time.Duration(time.Now().UnixNano() - parsed)

	started := time.Now()
	common.BotSession.ChannelMessageEdit(m.ChannelID, m.ID, "Gateway (http send -> gateway receive time): "+taken.String())
	httpPing := time.Since(started)

	common.BotSession.ChannelMessageEdit(m.ChannelID, m.ID, "HTTP API (Edit Msg): "+httpPing.String()+"\nGateway: "+taken.String())
}
