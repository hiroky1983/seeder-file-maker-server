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

func (p *Prompt) Prompt(c *openai.Client, prom string) string {
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
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}