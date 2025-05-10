package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"local.dev/DiscordBot/tasks"
	"local.dev/DiscordBot/tools"
)

func main() {
	if !tools.ReadTokenFromFile("token.json") {
		fmt.Println("Failed to read token file")
		return
	}

	// Create Discord session
	dg, err := discordgo.New("Bot " + tools.DiscordKey)
	if err != nil {
		fmt.Println("Failed to create Discord session:", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Failed to connect to Discord:", err)
		return
	}
	defer dg.Close()

	fmt.Println("Bot is now running.")

	onlineHi := "主人，我回来啦~"
	jt, err := tools.GetSoup(tools.JuheJitang)
	if err == nil {
		onlineHi += "\n" + jt
	}
	dg.ChannelMessageSend(tasks.ChannelGeneralID, onlineHi)

	tasks.InitTasks()

	dg.AddHandler(tasks.OnMessageCreate)

	// Start scheduled message routine
	tasks.InitDailyMessage()
	go tasks.StartDailyMessage(dg)

	// Handle graceful shutdown
	tools.WaitForQuit()
}
