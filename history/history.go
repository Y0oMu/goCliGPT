package history

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Conversation struct {
	History []Message `json:"history"`
}

// SaveConversation saves the conversation history to a JSON file with a timestamp
func SaveConversation(conversation *Conversation) {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("history_%s.json", timestamp)
	filepath := filepath.Join("history", filename)

	bytes, err := json.MarshalIndent(conversation, "", "  ")
	if err != nil {
		fmt.Println("Error saving conversation:", err)
		return
	}

	err = ioutil.WriteFile(filepath, bytes, 0644)
	if err != nil {
		fmt.Println("Error writing history file:", err)
	}
}

// ListConversations lists all saved conversation files
func ListConversations() ([]string, error) {
	var files []string

	err := filepath.Walk("history", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			files = append(files, info.Name())
		}
		return nil
	})

	return files, err
}

// LoadConversation loads a specific conversation history from a JSON file
func LoadConversation(filename string) (*Conversation, error) {
	filepath := filepath.Join("history", filename)

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening history file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading history file: %v", err)
	}

	var conversation Conversation
	err = json.Unmarshal(bytes, &conversation)
	if err != nil {
		return nil, fmt.Errorf("error parsing history file: %v", err)
	}

	return &conversation, nil
}
