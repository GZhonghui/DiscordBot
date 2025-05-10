package tasks

import (
	"github.com/bwmarrin/discordgo"
	"local.dev/DiscordBot/tools"
)

const ChannelLaughID = "1370664034675986503" // #laugh

func StartDailyJoke(s *discordgo.Session) {
	sendJoke := func() {
		var msg string = "哎呀，我想不到讲什么笑话给主人了"
		xh, xherr := tools.GetRandomJoke(tools.JuheXiaohua)
		if xherr == nil {
			for _, joke := range xh {
				msg = joke.Content + "\n亲爱的要开心哦~(ˊo̴̶̷̤ ̫ o̴̶̷̤ˋ)"
				break
			}
		}
		s.ChannelMessageSend(ChannelLaughID, msg)
	}

	sendJoke()
	f := tools.CreateDailyRandomTrigger(9, 21, 5, func() {
		sendJoke()
	}, tokyoLoc)

	go f()
}
