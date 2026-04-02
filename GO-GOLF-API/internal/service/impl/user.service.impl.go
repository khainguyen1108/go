package impl

import (
	"GO-GOLF-API/internal/models"
	"GO-GOLF-API/internal/repo"
	"GO-GOLF-API/internal/service"
	"GO-GOLF-API/pkg/response"
	JwtUtils "GO-GOLF-API/pkg/utils"
	PasswordEncoder "GO-GOLF-API/pkg/utils"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo repo.IUserRepository
}

// Register implements [service.IUserService].
func (us *UserService) Register(username string, purpose string) int {
	panic("unimplemented")
}

// GetUserInfoById implements [service.IUserService].
func (us *UserService) GetUserInfoById(id int) (*models.User, error) {
	user, err := us.userRepo.GetUserById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.NewAppError(http.StatusNotFound, response.ErrUserNotFound, err)
		}

		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrInternalError, err)
	}

	return user, nil
}

// Login implements [IUserService].
func (us *UserService) Login(loginUser models.LoginRequest) (*models.LoginResponse, interface{}) {
	user, err := us.userRepo.GetUserByUserId(loginUser.UserId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, response.NewAppError(http.StatusNotFound, response.ErrUserNotFound, err)
		}

		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrInternalError, err)
	}

	if user.UseStatus != 1 {
		return nil, response.NewAppError(http.StatusForbidden, response.ErrUserAlreadyUnregistered, nil)
	}

	if !PasswordEncoder.PasswordEncoderInstance.Verify(user.Password, loginUser.Password) {
		return nil, response.NewAppError(http.StatusUnauthorized, response.ErrPasswordNotMatch, nil)
	}

	return us.loginSuccessData(user, loginUser.UserAgent)
}

func NewUserService(userRepo repo.IUserRepository) service.IUserService {
	return &UserService{userRepo: userRepo}
}

// private functions
func (us *UserService) loginSuccessData(user *models.User, userAgent string) (*models.LoginResponse, interface{}) {
	sectionId := uuid.New().String()
	tokens, expiredAt, err := JwtUtils.GenerateTokens(user.Id, sectionId)

	if err != nil {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrGenerateTokenFailed, err)
	}

	_, err = us.userRepo.CreateAccountSession(&models.AccountSession{
		UserId:       user.Id,
		SessionId:    sectionId,
		RefreshToken: tokens.RefreshToken,
		DeviceInfo:   userAgent,
		ExpiresAt:    time.Unix(expiredAt, 0),
	})

	if err != nil {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrCreateAccountSessionFailed, err)
	}

	return tokens, nil
}
