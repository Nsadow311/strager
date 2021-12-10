package twitch

import (
	"fmt"

	"github.com/Nsadow311/stranger/common"
	"github.com/mediocregopher/radix/v3"
)

func (p *Plugin) Status() (string, string) {
	var unique int
	common.RedisPool.Do(radix.Cmd(&unique, "ZCARD", RedisKeyWebSubChannels))

	var numChannels int
	common.GORM.Model(&ChannelSubscription{}).Count(&numChannels)

<<<<<<< HEAD
	return "Twitch", fmt.Sprintf("%d/%d", unique, numChannels)
=======
	return "Youtube", fmt.Sprintf("%d/%d", unique, numChannels)
>>>>>>> 62f4d9c1b3ebfca2c67a21fce5befcac21acf4d7
}
