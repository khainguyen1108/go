package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc UserController) GetUserById(c *gin.Context) {
	name := c.DefaultQuery("name", "anonystick")
	uuid := c.Query("uuid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong.hhhh..ping" + name,
		"uuid":    uuid,
		"users":   []string{"cr7", "m10", "anonysitck"},
	})
}
