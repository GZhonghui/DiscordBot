package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type TokenConfig struct {
	DiscordKey string `json:"discord"`
	OpenaiKey  string `json:"openai"`
}

var (
	DiscordKey string
	OpenaiKey  string
)

func ReadTokenFromFile(path string) bool {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Read file failed:", err)
		return false
	}

	var config TokenConfig
	if err := json.Unmarshal(data, &config); err != nil {
		fmt.Println("Unpack JSON failed:", err)
		return false
	}

	DiscordKey = config.DiscordKey
	OpenaiKey = config.OpenaiKey
	return true
}

func WaitForQuit() {
	fmt.Println("Type 'q' then press Enter to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			fmt.Println("Quitting.")
			break
		}
	}
}
