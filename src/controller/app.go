package controller

import (
	"seeder-app/usecase"

	"github.com/sashabaranov/go-openai"
)

type App struct {
	OpenAiAPIKey *openai.Client
	PromptUseCase usecase.PromptUseCase
}

func NewApp(OpenAiAPIKey *openai.Client, PromptUseCase usecase.PromptUseCase) *App {
	return &App{
		OpenAiAPIKey: OpenAiAPIKey,
		PromptUseCase: PromptUseCase,
	}
}