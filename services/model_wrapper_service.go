package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
	"poker-degen/models"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func ModelWrapperService(req models.ModelRequest) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY cannot be empty")
	}

	url := "https://api.openai.com/v1/chat/completions"

	board := req.Board

	userPrompt := "
    Given the following, provide me the highest EV action and why:
    pot: 6bb
    community cards: "+board.CC1+", "+board.CC2+", "+board.CC3+"
    hero (BTN): "+board.HC1+", "+board.HC2+"
    villain (BB): "+board.V1C1+", "+board.V1C2+"
    flop: villain check, hero bet 2bb, villain raise 7bb
  "

	reqBody := ChatRequest{
		Model: "gpt-4.1-mini",
		Messages: []ChatMessage{
			{Role: "user", Content: userPrompt},
		},
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		panic(err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("StatusCode "+fmt.Sprintf("%d", resp.StatusCode))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		panic(err)
	}
	return chatResp.Choices[0].Message.Content, nil
}
