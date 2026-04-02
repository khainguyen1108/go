package repo

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/internal/models"
	"database/sql"
)

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserByUserId(userId string) (*models.User, error)
	CreateAccountSession(accountSession *models.AccountSession) (sql.Result, error)
	GetUserById(id int) (*models.User, error)
}

type UserRepository struct{}

// GetUserById implements [IUserRepository].
func (u *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User
	err := global.Mdb.Get(&user, "SELECT * FROM user WHERE id = ?", id)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateAccountSession implements [IUserRepository].
func (u *UserRepository) CreateAccountSession(accountSession *models.AccountSession) (sql.Result, error) {
	res, err := global.Mdb.NamedExec(`INSERT INTO account_sessions (user_id, session_id, refresh_token, device_info, expires_at) 
	VALUES (:user_id, :session_id, :refresh_token, :device_info, :expires_at)`, accountSession)

	if err != nil {
		return nil, err
	}

	return res, nil
}

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
