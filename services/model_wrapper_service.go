package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
	"strings"
	"pokerdegen/models"
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

func ModelWrapperService(req models.ModelRequest) ([]string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return []string{}, fmt.Errorf("OPENAI_API_KEY cannot be empty")
	}

	url := "https://api.openai.com/v1/chat/completions"

	board := req.Board


/**

** FINISHED **

Given the following, provide me the highest EV action and why:
pot: _
community cards: _, _, _, _, _
hero (BTN): _, _
villain (BB, _): _, _
flop: villain _, hero _, villain _,
turn: villain _, hero _, villain _
river: villain _, hero _, villain _
Respond in the format: [Check/Bet <amount>/Fold],[Villain's range as in integer in 0-100],[Hero's range as in integer in 0-100],[Brief explanation]
*/

/**

** WORKING **

Given the following, provide me the highest EV action and why:
pot: _
community cards: _, _, _,
hero (BTN): _, _
villain (BB): _, _
flop: villain _, hero _, villain _,
Respond in the format: [Check/Bet <amount>/Fold],[Villain's range as in integer in 0-100],[Hero's range as in integer in 0-100],[Brief explanation]
*/

	userPrompt := fmt.Sprintf(`
	Given the following, provide me the highest EV action and why:
	pot: %s
	community cards: %s, %s, %s
	hero (BTN): %s, %s
	villain (BB): %s, %s
	flop: villain check
	Respond in the format: [Check/Bet/Fold]; [Villain's range 0.0-1.0]; [Hero's range 0.0-1.0]; [Brief explanation]
	`, board.POT, board.CC1, board.CC2, board.CC3, board.HC1, board.HC2, board.V1C1, board.V1C2)

	reqBody := ChatRequest{
		Model: "gpt-4o",
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
		return []string{}, fmt.Errorf("StatusCode "+fmt.Sprintf("%d", resp.StatusCode))
	}

	var chatResp ChatResponse
	if err := json.Unmarshal(respBody, &chatResp); err != nil {
		panic(err)
	}
	return strings.Split(chatResp.Choices[0].Message.Content, ";"), nil
}
