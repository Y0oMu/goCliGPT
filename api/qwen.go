package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"goCliGPT/config"
)

type QwenApiRequest struct {
	Model      string       `json:"model"`
	Input      QwenInput    `json:"input"`
	Parameters QwenParams   `json:"parameters"`
}

type QwenInput struct {
	Messages []QwenMessage `json:"messages"`
}

type QwenMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type QwenParams struct {
	ResultFormat string `json:"result_format"`
}

type QwenApiResponse struct {
	Output struct {
		Choices []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	} `json:"output"`
}

// CallQwenApi function sends input to the Qwen API and gets the response
func CallQwenApi(history string) string {
	url := config.Config.Qwen.ApiUrl
	method := "POST"

	requestBody := QwenApiRequest{
		Model: config.Config.Qwen.ModelName,
		Input: QwenInput{
			Messages: []QwenMessage{
				{Role: "system", Content: "You are a helpful assistant."},
			},
		},
		Parameters: QwenParams{
			ResultFormat: "message",
		},
	}

	var historyMessages []QwenMessage
	err := json.Unmarshal([]byte(history), &historyMessages)
	if err != nil {
		fmt.Println("History unmarshal error:", err)
		return "Error in parsing history"
	}
	requestBody.Input.Messages = append(requestBody.Input.Messages, historyMessages...)

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
	req.Header.Add("Authorization", "Bearer "+config.Config.Qwen.ApiKey)
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

	var apiResponse QwenApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return "Error in parsing response"
	}

	if len(apiResponse.Output.Choices) > 0 && len(apiResponse.Output.Choices[0].Message.Content) > 0 {
		return apiResponse.Output.Choices[0].Message.Content
	}
	return "No response"
}
