package controller

import (
	"GO-ECOMMERCE-BACKEND-API/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, response.CodeSuccess, map[string]string{"id": c.Param("id")})
}
