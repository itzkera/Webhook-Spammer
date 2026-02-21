package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func sendMessage(url, text string) {
	msg := Message{Content: text}
	payload, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshaling message:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Message sent:", text)
}

func main() {
	webhookURL := ""

	for i := 1; i <= 1000; i++ {
		text := fmt.Sprintf("wow %d", i)
		sendMessage(webhookURL, text)
	}
}