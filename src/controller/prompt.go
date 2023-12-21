package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) Prompt(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	prmpt, err := a.PromptUseCase.Prompt(a.OpenAiAPIKey, string(bodyBytes))
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"prompt": prmpt,
	})
}
