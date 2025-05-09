package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadTokenFromFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read token file:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(data))
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
