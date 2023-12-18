package repository

import "github.com/sashabaranov/go-openai"

//go:generate moq -pkg fakerepository -out ../../moq/fakerepository/prompt_test_moq.go . Prompt

type Prompt interface {
	Prompt(*openai.Client, string) (string, error)
}