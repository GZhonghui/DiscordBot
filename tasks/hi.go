package tasks

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"

	"local.dev/DiscordBot/tools"
)

const (
	ChannelHiID = "1370102028684099714" // #hi
	TargetHour  = 8
	TargetMin   = 5
)

var StartDate time.Time

func InitDailyMessage() {
	StartDate = time.Date(2025, 5, 9, 0, 0, 0, 0, tokyoLoc)
}

func StartDailyMessage(s *discordgo.Session) {
	for {
		now := time.Now().In(tokyoLoc)
		next := time.Date(now.Year(), now.Month(), now.Day(), TargetHour, TargetMin, 0, 0, tokyoLoc)
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}
		sleepDuration := next.Sub(now)
		fmt.Printf("Waiting until %v to send message...\n", next.Format("15:04:05"))
		time.Sleep(sleepDuration)

		// Construct message
		today := time.Now().In(tokyoLoc)
		daysSince := int(today.Sub(StartDate).Hours()/24) + 1
		year, month, day := today.Date()
		msg := fmt.Sprintf("早上好、亲爱的！(/≧▽≦/)\n今天是%d年%d月%d日、是爱你第%d天哦~\n今天也要好好爱自己哦(◕ˇ∀ˇ◕。)", year, int(month), day, daysSince)

		jt, jterr := tools.GetSoup(tools.JuheJitang)
		if jterr == nil {
			msg += "\n" + jt
		}

		_, err := s.ChannelMessageSend(ChannelHiID, msg)
		if err != nil {
			fmt.Println("Failed to send message:", err)
		} else {
			// fmt.Println("Message sent:", msg)
		}
	}
}
