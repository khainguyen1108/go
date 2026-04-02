package service

import (
	"GO-GOLF-API/internal/models"
)

type IUserService interface {
	Register(username, purpose string) int
	Login(loginUser models.LoginRequest) (*models.LoginResponse, interface{})
	GetUserInfoById(id int) (*models.User, error)
}
