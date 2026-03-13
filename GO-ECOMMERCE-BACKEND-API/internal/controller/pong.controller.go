package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
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
	client := openai.NewClient(
		option.WithAPIKey("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"),
	)

	resp, err := client.Responses.New(context.TODO(), responses.ResponseNewParams{
		Model: "gpt-5.4",
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String("Say this is a test")},
	})
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"response": resp.OutputText()})
}
