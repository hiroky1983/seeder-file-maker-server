package usecase

import (
	"seeder-app/domain/model/prompt"
	"seeder-app/domain/repository"

	"github.com/sashabaranov/go-openai"
)
type PromptUseCase interface {
	Prompt(*openai.Client, string) string
}

type Prompt struct {
	promptRepo repository.Prompt
}

func NewPrompt(promptRepo repository.Prompt) *Prompt {
	return &Prompt{
		promptRepo: promptRepo,
	}
}

func (p *Prompt) Prompt(c *openai.Client, prom string) string {
	template := prompt.GeneratePromptText(prom)
	prompt := p.promptRepo.Prompt(c, template)
	return prompt
}
