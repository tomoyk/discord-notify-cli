package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const maxContentLength = 2000 - 6 // ``` and ```

func main() {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		fmt.Fprintf(os.Stderr, "Error: DISCORD_WEBHOOK_URL environment variable is not set\n")
		os.Exit(1)
	}

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}

	content := string(input)
	chunks := splitContent(content, maxContentLength)

	for _, chunk := range chunks {
		payload := map[string]string{
			"content": "```" + chunk + "```",
		}

		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding payload to JSON: %v\n", err)
			os.Exit(1)
		}

		resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadBytes))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error sending request to Discord: %v\n", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Fprintf(os.Stderr, "Discord webhook returned error: %s\n", string(body))
			os.Exit(1)
		}
	}

	fmt.Println("Notification sent to Discord successfully.")
}

func splitContent(content string, length int) []string {
	var chunks []string
	for len(content) > length {
		chunks = append(chunks, content[:length])
		content = content[length:]
	}
	if len(content) > 0 {
		chunks = append(chunks, content)
	}
	return chunks
}
