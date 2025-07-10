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

	userPrompt := "What is the capital of France?"

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

	// Use a different name to avoid overwriting the original `req`
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
