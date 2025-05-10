package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// GetSoup retrieves a motivational quote from the Juhe API
func GetSoup(apiKey string) (string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	url := "https://apis.juhe.cn/fapig/soup/query?key=" + apiKey

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result struct {
		Reason    string `json:"reason"`
		ErrorCode int    `json:"error_code"`
		Result    struct {
			Text string `json:"text"`
		} `json:"result"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	if result.ErrorCode != 0 {
		return "", errors.New("API returned error: " + result.Reason)
	}

	return result.Result.Text, nil
}
