package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	ChannelID  = "1370102028684099714" // Replace with your Discord channel ID
	TargetHour = 8
	TargetMin  = 5
)

var (
	tokyoLoc  *time.Location
	StartDate time.Time
)

func main() {
	tokyoLoc, _ = time.LoadLocation("Asia/Tokyo")
	StartDate = time.Date(2025, 5, 9, 0, 0, 0, 0, tokyoLoc)

	token := readTokenFromFile("token.txt")

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

	// Start scheduled message routine
	go startDailyMessage(dg, ChannelID)

	// Handle graceful shutdown
	waitForQuit()
}

func readTokenFromFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read token file:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(data))
}

func startDailyMessage(s *discordgo.Session, channelID string) {
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

		_, err := s.ChannelMessageSend(channelID, msg)
		if err != nil {
			fmt.Println("Failed to send message:", err)
		} else {
			// fmt.Println("Message sent:", msg)
		}
	}
}

func waitForQuit() {
	fmt.Println("Type 'q' then press Enter to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			fmt.Println("Quitting.")
			break
		}
	}
}
