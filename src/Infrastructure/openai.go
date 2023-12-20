package infrastructure

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Prompt struct{}

func NewPrompt() *Prompt {
	return &Prompt{}
}

const (
	maxTokens = 150
)

func (p *Prompt) Prompt(c *openai.Client, prom string) (string, error) {
	resp, err := c.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prom,
				},
			},
			Temperature:0.0,
			MaxTokens:   maxTokens,
		},
	)

	if err != nil {
		return "", err
	}

	fmt.Printf("Response: %+v\n", resp)
	return resp.Choices[0].Message.Content, nil
}