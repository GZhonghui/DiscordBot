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

	onlineHi := "主人，我回来啦~" + tools.RandomFace()
	jt, jterr := tools.GetSoup(tools.JuheJitang)
	if jterr == nil {
		onlineHi += "\n" + jt
	}
	dg.ChannelMessageSend(tasks.ChannelGeneralID, onlineHi)

	tasks.InitTasks()

	dg.AddHandler(tasks.OnMessageCreate)

	// Start scheduled message routine
	tasks.InitDailyMessage()
	go tasks.StartDailyMessage(dg)

	tasks.StartDailyJoke(dg)

	// Handle graceful shutdown
	tools.WaitForQuit()
}
