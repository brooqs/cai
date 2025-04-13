package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"cai/config"

	openai "github.com/sashabaranov/go-openai"
)

func cleanResponse(response string) string {
	// Remove code blocks and unnecessary markdown
	re := regexp.MustCompile("(?s)```.*?\\n(.*?)```")
	matches := re.FindStringSubmatch(response)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	// Remove plaintext or similar unwanted terms
	response = strings.ReplaceAll(response, "plaintext", "")
	return strings.TrimSpace(response)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cai \"your query here\"")
		return
	}

	prompt := strings.Join(os.Args[1:], " ")
	promptFull := fmt.Sprintf("Only provide the command or configuration without explanations: \"%s\"", prompt)

	apiKey, model, maxTokens := config.LoadAPIConfig()

	client := openai.NewClient(apiKey)

	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: promptFull,
			},
		},
		MaxTokens: maxTokens,
	})

	if err != nil {
		log.Fatalf("OpenAI request error: %v", err)
	}

	cleanedResponse := cleanResponse(resp.Choices[0].Message.Content)
	fmt.Println(cleanedResponse)
}
