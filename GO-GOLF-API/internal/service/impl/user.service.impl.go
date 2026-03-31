package impl

import (
	"GO-GOLF-API/internal/models"
	"GO-GOLF-API/internal/repo"
	"GO-GOLF-API/internal/service"
)

type UserService struct {
	userRepo repo.IUserRepository
}

// Login implements [IUserService].
func (us *UserService) Login(loginUser models.LoginRequest) (*models.User, error) {
	return us.userRepo.GetUserByUserId(loginUser.UserId)
}

// Register implements [IUserService].
func (us *UserService) Register(username string, purpose string) int {
	panic("unimplemented")
}

func NewUserService(userRepo repo.IUserRepository) service.IUserService {
	return &UserService{userRepo: userRepo}
}
