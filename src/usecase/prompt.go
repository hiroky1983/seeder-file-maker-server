package usecase

import (
	"fmt"
	"seeder-app/domain/model/prompt"
	"seeder-app/domain/repository"

	"github.com/sashabaranov/go-openai"
)

type PromptUseCase interface {
	Prompt(*openai.Client, string) (string, error)
}

type Prompt struct {
	promptRepo repository.Prompt
}

func NewPrompt(promptRepo repository.Prompt) *Prompt {
	return &Prompt{
		promptRepo: promptRepo,
	}
}

func (p *Prompt) Prompt(c *openai.Client, prom string) (string, error) {
	// if err := prompt.ValidatePromptText(prom); err != nil {
	// 	return "", err
	// }

	template := prompt.GeneratePromptText(prom)
	prompt, err := p.promptRepo.Prompt(c, template)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return prompt, nil
}
