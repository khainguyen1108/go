package repo

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/internal/models"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserByUserId(userId string) (*models.User, error)
}

type UserRepository struct{}

// GetUserByUserId implements [IUserRepository].
func (u *UserRepository) GetUserByUserId(userId string) (*models.User, error) {
	var user models.User
	err := global.Mdb.Get(&user, "SELECT * FROM user WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail implements [IUserRepository].
func (u *UserRepository) GetUserByEmail(email string) bool {
	panic("unimplemented")
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}
