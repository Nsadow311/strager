package allocstat

import (
	"fmt"
	"runtime"
	"time"

	"github.com/Nsadow311/stranger/commands"
	"github.com/Nsadow311/stranger/common"
	"github.com/Nsadow311/stranger/stdcommands/util"
	"github.com/jonas747/dcmd/v4"
)

var Command = &commands.YAGCommand{
	CmdCategory:          commands.CategoryDebug,
	HideFromCommandsPage: true,
	Name:                 "allocstat",
	Description:          "Memory statistics.",
	HideFromHelp:         true,
	RunFunc: util.RequireBotAdmin(func(data *dcmd.Data) (interface{}, error) {
		common.BotSession.ChannelTyping(data.ChannelID)
		var memstatsStarted runtime.MemStats
		runtime.ReadMemStats(&memstatsStarted)

		time.Sleep(time.Second * 10)

		var memstatsStopped runtime.MemStats
		runtime.ReadMemStats(&memstatsStopped)

		bytesAlloc := (memstatsStopped.TotalAlloc - memstatsStarted.TotalAlloc) / 1000
		numAlloc := memstatsStopped.Mallocs - memstatsStarted.Mallocs

		lastGC := time.Unix(0, int64(memstatsStopped.LastGC))
		numGC := memstatsStopped.NumGC

		return fmt.Sprintf("Bytes allocated(10s): %dKB\nNum allocs (10s): %d\nLast gc: %s\nNum gc (from start): %d", bytesAlloc, numAlloc, time.Since(lastGC), numGC), nil
	}),
}
