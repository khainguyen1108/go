package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/genai"
)

type TestController struct{}

func (tc TestController) Pong(c *gin.Context) {

	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text("generate 1 video for me"),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"response": result.Text()})
}

func (tc TestController) PongOpenApi(c *gin.Context) {

}
