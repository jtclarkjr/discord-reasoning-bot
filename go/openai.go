package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// Initialize OpenAI client
func initOpenAIClient(apiKey string) *openai.Client {
	return openai.NewClient(option.WithAPIKey(apiKey))
}

// Check if a message is offensive using OpenAI's API
func isMessageOffensive(client *openai.Client, message string) (bool, error) {
	ctx := context.Background()

	// Define the prompt for the model
	prompt := fmt.Sprintf("Is the following message offensive? Answer with 'true' or 'false' and no period:\n\n\"%s\"", message)

	// Create a chat completion request
	chatCompletion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Model: openai.F(openai.ChatModelO3Mini), // Check OpenAI account has tier access to o3-mini model if getting error
	})
	if err != nil {
		return false, err
	}

	// Check the model's response
	if len(chatCompletion.Choices) > 0 {
		answer := chatCompletion.Choices[0].Message.Content
		if strings.ToLower(answer) == "true" {
			return true, nil
		}
	}

	return false, nil
}
