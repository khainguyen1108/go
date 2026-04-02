package controller

import (
	"GO-GOLF-API/internal/models"
	"GO-GOLF-API/internal/service"
	"GO-GOLF-API/pkg/response"
	"GO-GOLF-API/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUserById(c *gin.Context) (interface{}, error) {
	user := c.MustGet("user")
	return user, nil
}

func (uc *UserController) Login(c *gin.Context) (interface{}, error) {
	// Add login logic here return one map
	var loginUser models.LoginRequest

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		return nil, response.NewAppError(http.StatusBadRequest, response.ErrValidationFailed, err)
	}
	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrInternalError, gin.Error{})
	}

	if apiErr := utils.ValidateStruct(loginUser, validation.(*validator.Validate)); apiErr != nil {
		return nil, apiErr
	}

	loginUser.UserAgent = c.Request.UserAgent()
	responseData, err := uc.userService.Login(loginUser)

	if err != nil {
		return nil, err.(*response.AppError)
	}

	return responseData, nil
}

func (uc *UserController) LogOut(c *gin.Context) (interface{}, error) {
	sessionId := c.GetString("section_id")
	if sessionId == "" {
		return nil, response.NewAppError(http.StatusBadRequest, response.ErrLogoutFailed, gin.Error{})
	}

	return nil, uc.userService.LogOut(sessionId)
}

func (uc *UserController) Refresh(c *gin.Context) (interface{}, error) {
	var refreshRequest models.RefreshRequest

	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		return nil, response.NewAppError(http.StatusBadRequest, response.ErrValidationFailed, err)
	}
	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrInternalError, gin.Error{})
	}

	if apiErr := utils.ValidateStruct(refreshRequest, validation.(*validator.Validate)); apiErr != nil {
		return nil, apiErr
	}
	refreshRequest.UserAgent = c.Request.UserAgent()
	responseData, err := uc.userService.Refresh(refreshRequest)

	if err != nil {
		return nil, err.(*response.AppError)
	}

	return responseData, nil
}
