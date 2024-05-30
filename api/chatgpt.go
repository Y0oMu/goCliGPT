package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"goCliGPT/config"
)

type ChatGptApiResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// CallChatGptApi function sends input to the ChatGPT API and gets the response
func CallChatGptApi(input string) string {
	url := config.Config.ChatGPT.ApiUrl
	method := "POST"

	var jsonStr = []byte(`{
		"prompt": "` + input + `",
		"max_tokens": 100
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Request error:", err)
		return ""
	}
	req.Header.Add("Authorization", "Bearer " + config.Config.ChatGPT.ApiKey)
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

	var apiResponse ChatGptApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("JSON unmarshal error:", err)
		return "Error in parsing response"
	}

	if len(apiResponse.Choices) > 0 {
		return apiResponse.Choices[0].Text
	}
	return "No response"
}
