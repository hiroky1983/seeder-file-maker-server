package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func(a *App) Prompt(c *gin.Context)  {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	prmpt := a.PromptUseCase.Prompt(a.OpenAiAPIKey, string(bodyBytes))
	c.JSON(200, gin.H{
		"prompt": prmpt,
	})
}
