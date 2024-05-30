package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io"

	"goCliGPT/api"
	"goCliGPT/config"
)

func main() {
	config.LoadConfig()

	info, _ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo 'Your Message qwen' | ./goCliGPT")
		return
	}

	var lines []string
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	if len(lines) > 0 && strings.HasSuffix(lines[0], "qwen") {
		response := api.CallQwenApi(lines[0])
		fmt.Println(response)
		return
	}

	if len(lines) > 0 && strings.HasSuffix(lines[0], "chatgpt") {
		response := api.CallChatGptApi(lines[0])
		fmt.Println(response)
		return
	}

	fmt.Println()
}
