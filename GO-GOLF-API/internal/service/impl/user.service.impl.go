package impl

import (
	"GO-GOLF-API/internal/models"
	"GO-GOLF-API/internal/repo"
	"GO-GOLF-API/internal/service"
	"GO-GOLF-API/pkg/response"
	JwtUtils "GO-GOLF-API/pkg/utils"
	PasswordEncoder "GO-GOLF-API/pkg/utils"
	RedisUtils "GO-GOLF-API/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo repo.IUserRepository
}

// Refresh implements [service.IUserService].
func (us *UserService) Refresh(refreshRequest models.RefreshRequest) (*models.LoginResponse, error) {

	claims, err := JwtUtils.VerifyAndParseJWT(refreshRequest.RefreshToken)

	if err != nil {
		return nil, response.NewAppError(http.StatusUnauthorized, response.ErrInvalidRefreshToken, err)
	}

	sessionId := claims["section_id"].(string)
	userId := int(claims["user_id"].(float64))
	//check refresh token hợp lệ (chưa bị sử dụng)
	accountSession, _ := us.userRepo.GetAccountSessionByRefreshToken(refreshRequest.RefreshToken, sessionId)

	if accountSession.IsUsed == 1 {
		_, errRevolked := us.userRepo.RevokedAccountSession(userId)
		if errRevolked != nil {
			return nil, response.NewAppError(http.StatusInternalServerError, response.ErrRevokeAccountFailed, errRevolked)
		}
		RedisUtils.Set(context.Background(), strconv.Itoa(userId), time.Now().Unix(), 15*time.Minute)
		return nil, response.NewAppError(http.StatusUnauthorized, response.ErrInvalidRefreshToken, errors.New("refresh token already used"))
	}
	//check refresh token hợp lệ (chưa bị thu hồi)

	if accountSession.IsRevoked == 1 {
		return nil, response.NewAppError(http.StatusUnauthorized, response.ErrInvalidRefreshToken, err)
	}

	tokens, generateErr := JwtUtils.GenerateTokens(userId, sessionId)

	if generateErr != nil {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrGenerateTokenFailed, generateErr)
	}

	errCreateAccounSession := us.userRepo.RotateAccountSession(sessionId, &models.AccountSession{
		UserId:       userId,
		SessionId:    sessionId,
		RefreshToken: tokens.RefreshToken,
		DeviceInfo:   refreshRequest.UserAgent,
	})

	if errCreateAccounSession != nil {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrCreateAccountSessionFailed, errCreateAccounSession)
	}

	return tokens, nil
}

// LogOut implements [service.IUserService].
func (us *UserService) LogOut(sessionId string) error {
	_, err := us.userRepo.UpdateAccountSession(sessionId)
	if err != nil {
		return response.NewAppError(http.StatusInternalServerError, response.ErrLogoutFailed, err)
	}
	return nil
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
	tokens, err := JwtUtils.GenerateTokens(user.Id, sectionId)

	if err != nil {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrGenerateTokenFailed, err)
	}

	_, err = us.userRepo.CreateAccountSession(&models.AccountSession{
		UserId:       user.Id,
		SessionId:    sectionId,
		RefreshToken: tokens.RefreshToken,
		DeviceInfo:   userAgent,
	})

	if err != nil {
		return nil, response.NewAppError(http.StatusInternalServerError, response.ErrCreateAccountSessionFailed, err)
	}

	return tokens, nil
}
