package main

import (
	"context"
	"fmt"
	infrastructure "seeder-app/Infrastructure"
	"seeder-app/config"
	"seeder-app/controller"
	"seeder-app/router"
	"seeder-app/usecase"

	openai "github.com/sashabaranov/go-openai"
)

func initOpenAi (c *config.Config) *openai.Client {
	fmt.Println(c.OpenAiAPIKey)
	client := openai.NewClient(c.OpenAiAPIKey)
	fmt.Println(client)
	return client

}

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		fmt.Println(err)
	}

	o := initOpenAi(cfg)
	api := controller.NewApp(
		o,
		usecase.NewPrompt(infrastructure.NewPrompt()),
	)
	
	r := router.NewRouter(api)
	
  r.Run() 
}
