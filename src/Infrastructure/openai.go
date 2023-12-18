package infrastructure

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Prompt struct{}

func NewPrompt() *Prompt {
	return &Prompt{}
}

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
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}