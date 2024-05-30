package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"goCliGPT/config"
)

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
func CallQwenApi(input string) string {
	url := config.Config.Qwen.ApiUrl
	method := "POST"

	var jsonStr = []byte(`{
		"model": "` + config.Config.Qwen.ModelName + `",
		"input": {
			"messages": [
				{"role": "system", "content": "You are a helpful assistant."},
				{"role": "user", "content": "` + input + `"}
			]
		},
		"parameters": {
			"result_format": "message"
		}
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Request error:", err)
		return ""
	}
	req.Header.Add("Authorization", "Bearer " + config.Config.Qwen.ApiKey)
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
