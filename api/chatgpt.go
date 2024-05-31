package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"goCliGPT/config"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGptApiRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type ChatGptApiResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

// CallChatGptApi function sends input to the ChatGPT API and gets the response
func CallChatGptApi(history string) string {
	url := config.Config.ChatGPT.ApiUrl
	method := "POST"

	var historyMessages []Message
	err := json.Unmarshal([]byte(history), &historyMessages)
	if err != nil {
		fmt.Println("History unmarshal error:", err)
		return "Error in parsing history"
	}

	requestBody := ChatGptApiRequest{
		Model: config.Config.ChatGPT.ModelName,
		Messages: append([]Message{
			{Role: "system", Content: "You are a helpful assistant."},
		}, historyMessages...),
		MaxTokens: 100,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("JSON marshal error:", err)
		return "Error in creating request body"
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Request error:", err)
		return ""
	}
	req.Header.Add("Authorization", "Bearer "+config.Config.ChatGPT.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("API call error:", err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read body error:", err)
		return ""
	}

	//fmt.Println("API Response:", string(body)) // Add this line for debugging

	var apiResponse ChatGptApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return "Error in parsing response"
	}

	if len(apiResponse.Choices) > 0 && len(apiResponse.Choices[0].Message.Content) > 0 {
		return apiResponse.Choices[0].Message.Content
	}
	return "No response"
}
