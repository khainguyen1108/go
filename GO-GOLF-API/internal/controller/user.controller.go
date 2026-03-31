package controller

import (
	"GO-GOLF-API/internal/models"
	"GO-GOLF-API/internal/service"
	"GO-GOLF-API/pkg/response"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, response.CodeSuccess, map[string]string{"id": c.Param("id")})
}

func (uc *UserController) Login(c *gin.Context) {
	// Add login logic here return one map
	var loginUser models.LoginRequest

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.Error(&response.AppError{
			StatusCode: http.StatusBadRequest,
			AppCode:    response.ErrValidationFailed,
			RootErr:    err,
		})
		return
	}

	responseData, err := uc.userService.Login(loginUser)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.ErrorResponse(c, http.StatusNotFound, response.ErrUserNotFound)
			return
		}

		c.Error(&response.AppError{
			StatusCode: http.StatusInternalServerError,
			AppCode:    response.ErrInternalError,
			RootErr:    err,
		})
		return
	}

	response.SuccessResponse(c, response.CodeSuccess, responseData)
}
