package service

import (
	"GO-GOLF-API/internal/models"
)

type IUserService interface {
	Register(username, purpose string) int
	Login(loginUser models.LoginRequest) (*models.User, error)
}
