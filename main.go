package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"local.dev/DiscordBot/tasks"
	"local.dev/DiscordBot/tools"
)

func main() {
	token := tools.ReadTokenFromFile("token.txt")

	// Create Discord session
	dg, err := discordgo.New("Bot " + token)
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

	tasks.InitTasks()

	// Start scheduled message routine
	tasks.InitDailyMessage()
	go tasks.StartDailyMessage(dg)

	// Handle graceful shutdown
	tools.WaitForQuit()
}
