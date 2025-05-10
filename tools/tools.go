package tools

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type TokenConfig struct {
	DiscordKey  string `json:"discord"`
	OpenaiKey   string `json:"openai"`
	JuheJitang  string `json:"jhsj_xljt"`
	JuheXiaohua string `json:"jhsj_xh"`
}

var (
	DiscordKey  string
	OpenaiKey   string
	JuheJitang  string
	JuheXiaohua string
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
	JuheJitang = config.JuheJitang
	JuheXiaohua = config.JuheXiaohua
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
