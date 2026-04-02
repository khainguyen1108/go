package repo

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/internal/models"
	"database/sql"
)

type SQLRunner interface {
	Exec(query string, args ...any) (sql.Result, error)
	NamedExec(query string, arg any) (sql.Result, error)
}

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserByUserId(userId string) (*models.User, error)
	GetUserById(id int) (*models.User, error)

	CreateAccountSession(accountSession *models.AccountSession) (sql.Result, error)
	UpdateAccountSession(sessionId string) (sql.Result, error)
	RotateAccountSession(oldSessionId string, newSession *models.AccountSession) error
	GetAccountSessionByRefreshToken(refreshToken string, sessionId string) (*models.AccountSession, error)
	RevokedAccountSession(userId int) (sql.Result, error)
}

type UserRepository struct{}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

// =========================================================================
// CÁC HÀM CORE (PRIVATE)
// =========================================================================

func (u *UserRepository) coreCreateAccountSession(runner SQLRunner, accountSession *models.AccountSession) (sql.Result, error) {
	query := `
		INSERT INTO account_sessions (user_id, session_id, refresh_token, device_info) 
		VALUES (:user_id, :session_id, :refresh_token, :device_info)
	`
	return runner.NamedExec(query, accountSession)
}

func (u *UserRepository) coreUpdateAccountSession(runner SQLRunner, sessionId string) (sql.Result, error) {
	query := `UPDATE account_sessions SET is_used = 1 WHERE session_id = ?`
	return runner.Exec(query, sessionId)
}

func (u *UserRepository) coreRevokedAccountSession(runner SQLRunner, userId int) (sql.Result, error) {
	query := `UPDATE account_sessions SET is_revoked = 1 WHERE user_id = ?`
	return runner.Exec(query, userId)
}

// =========================================================================
// CÁC HÀM PUBLIC DÀNH CHO TẦNG SERVICE GỌI
// =========================================================================

// CreateAccountSession (Chạy độc lập)
func (u *UserRepository) RevokedAccountSession(userId int) (sql.Result, error) {
	return u.coreRevokedAccountSession(global.Mdb, userId)
}

func (u *UserRepository) CreateAccountSession(accountSession *models.AccountSession) (sql.Result, error) {
	return u.coreCreateAccountSession(global.Mdb, accountSession)
}

// UpdateAccountSession (Chạy độc lập)
func (u *UserRepository) UpdateAccountSession(sessionId string) (sql.Result, error) {
	return u.coreUpdateAccountSession(global.Mdb, sessionId)
}

// RotateAccountSession (Chạy trong Transaction - Dành cho hàm Refresh của bạn)
func (u *UserRepository) RotateAccountSession(oldSessionId string, newSession *models.AccountSession) error {
	// Khởi tạo Transaction
	tx, err := global.Mdb.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if _, err := u.coreUpdateAccountSession(tx, oldSessionId); err != nil {
		return err
	}
	if _, err := u.coreCreateAccountSession(tx, newSession); err != nil {
		return err
	}
	return tx.Commit()
}

// =========================================================================
// CÁC HÀM ĐỌC DỮ LIỆU (SELECT) GIỮ NGUYÊN NHƯ CŨ
// =========================================================================

func (u *UserRepository) GetUserById(id int) (*models.User, error) {
	var user models.User
	err := global.Mdb.Get(&user, "SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetUserByUserId(userId string) (*models.User, error) {
	var user models.User
	err := global.Mdb.Get(&user, "SELECT * FROM user WHERE user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetAccountSessionByRefreshToken(refreshToken string, sessionId string) (*models.AccountSession, error) {
	var accountSession models.AccountSession
	err := global.Mdb.Get(&accountSession, "SELECT * FROM account_sessions WHERE refresh_token = ? AND session_id = ?", refreshToken, sessionId)
	if err != nil {
		return nil, err
	}
	return &accountSession, nil
}

func (u *UserRepository) GetUserByEmail(email string) bool {
	panic("unimplemented")
}
