package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"goCliGPT/api"
	"goCliGPT/config"
	"goCliGPT/history"
)

func main() {
	config.LoadConfig()
	fmt.Println("Usage: ./goCliGPT <api_name>")
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./goCliGPT <api_name>")
		return
	}

	mode := os.Args[1]

	// List available conversation histories
	files, err := history.ListConversations()
	if err != nil {
		fmt.Println("Error listing conversations:", err)
		return
	}

	fmt.Println("Available conversations:")
	for i, file := range files {
		fmt.Printf("%d: %s\n", i+1, file)
	}

	// Prompt user to select a conversation or start a new one
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of the conversation to load or 'new' to start a new conversation: ")
	selection, _ := reader.ReadString('\n')
	selection = strings.TrimSpace(selection)

	var conversation *history.Conversation

	if selection == "new" {
		conversation = &history.Conversation{}
	} else {
		index := -1
		fmt.Sscanf(selection, "%d", &index)
		if index > 0 && index <= len(files) {
			conversation, err = history.LoadConversation(files[index-1])
			if err != nil {
				fmt.Println("Error loading conversation:", err)
				return
			}
			fmt.Println("Loaded conversation history. Type 'new' to start a new conversation, or continue with the current one.")
		} else {
			fmt.Println("Invalid selection, starting a new conversation.")
			conversation = &history.Conversation{}
		}
	}

	// Main loop for conversation
	for {
		fmt.Print("You: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			history.SaveConversation(conversation)
			break
		} else if input == "new" {
			conversation = &history.Conversation{}
			continue
		}

		conversation.History = append(conversation.History, history.Message{Role: "user", Content: input})

		var response string
		if mode == "qwen" {
			response = api.CallQwenApi(formatInput(conversation.History))
		} else if mode == "chatgpt" {
			response = api.CallChatGptApi(formatInput(conversation.History))
		} else {
			fmt.Println("Unknown mode:", mode)
			continue
		}

		conversation.History = append(conversation.History, history.Message{Role: "assistant", Content: response})
		fmt.Println("Assistant:", response)
	}
}

func formatInput(history []history.Message) string {
	var formatted []string
	for _, msg := range history {
		formatted = append(formatted, fmt.Sprintf("{\"role\": \"%s\", \"content\": \"%s\"}", msg.Role, msg.Content))
	}
	result := "[" + strings.Join(formatted, ", ") + "]"
	//fmt.Println("Formatted Input:", result)
	return result
}
