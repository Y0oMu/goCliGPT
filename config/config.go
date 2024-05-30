package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ApiConfig struct {
	ApiUrl    string `json:"api_url"`
	ApiKey    string `json:"api_key"`
	ModelName string `json:"model_name"`
}

type Configuration struct {
	ChatGPT ApiConfig `json:"ChatGPT"`
	Qwen    ApiConfig `json:"Qwen"`
}

var Config Configuration

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return
	}
}
