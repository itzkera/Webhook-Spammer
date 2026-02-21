package main 

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"time"
)


type Message struct {
	Content string `json:"content"`
}


func message(url, text string) (*http.Response, error) {
	msg := Message{Content: text}
	payload, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	return resp, nil
}


func main() {
	webhookURL := "https://discord.com/api/webooks/"
	// ur webhook here
	for i := 0; i < 10; i++ {
	msg := Message{
    Content: fmt.Sprintf("@here %d", i+1),
}
payload, err := json.Marshal(msg)

		if err != nil {
			fmt.Printf("Error marshaling message: %v\n", err)
			continue
		}
		remaining := rasp.Header.Get("X-RateLimit-Remaining")
		resetAfter := rasp.Header.Get("X-RateLimit-Reset-After")
		fmt.Printf("Remaining: %s, Reset After: %s seconds\n", remaining, resetAfter)
		if remaining == "0" || resp.StatusCode == 429 {
			fmt.Printf("Rate limit hit. Waiting for %s seconds before retrying...\n", resetAfter)
				time.Sleep(1 * time.Second)
		defer resp.Body.Close()
		}
 }
}