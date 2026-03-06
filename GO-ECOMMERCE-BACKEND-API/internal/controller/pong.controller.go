package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct{}

func (tc TestController) Pong(c *gin.Context) {
	fmt.Printf("ClientIP: %s\n", c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
