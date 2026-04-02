package utils

import (
	"GO-GOLF-API/global"
	"GO-GOLF-API/internal/models"
	"GO-GOLF-API/pkg/response"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// password encoder utils

type PasswordEncoder interface {
	Encode(password string) (string, error)
	Verify(encodedPassword, rawPassword string) bool
}

type BcryptEncoder struct{}

// Verify implements [PasswordEncoder].
func (b *BcryptEncoder) Verify(encodedPassword string, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	return err == nil
}

func (b *BcryptEncoder) Encode(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func NewPasswordEncoder() PasswordEncoder {
	return &BcryptEncoder{}
}

var PasswordEncoderInstance = NewPasswordEncoder()

// jwt utils

var (
	accessTokenSecret  = []byte(global.Config.Jwt.AccessSecretKey)
	refreshTokenSecret = []byte(global.Config.Jwt.RefreshSecretKey)
)

func GenerateTokens(userId int, sectionId string) (*models.LoginResponse, int64, error) {
	accessClaims := jwt.MapClaims{
		"user_id":    userId,
		"section_id": sectionId,
		"exp":        time.Now().Add(15 * time.Minute).Unix(),
		"iat":        time.Now().Unix(),
	}

	accessTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := accessTokenObj.SignedString(accessTokenSecret)
	if err != nil {
		return nil, 0, err
	}
	expiredAt := time.Now().Add(24 * time.Hour).Unix()
	refreshClaims := jwt.MapClaims{
		"user_id":    userId,
		"section_id": sectionId,
		"exp":        expiredAt,
		"iat":        time.Now().Unix(),
	}

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenObj.SignedString(refreshTokenSecret)
	if err != nil {
		return nil, 0, err
	}

	return &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, expiredAt, nil
}

func VerifyAndParseJWT(tokenString string) (jwt.MapClaims, *response.AppError) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, response.NewAppError(http.StatusUnauthorized, response.ErrAlogrithmNotSupported, gin.Error{})
		}
		return []byte(accessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, response.NewAppError(http.StatusUnauthorized, response.ErrTokenExpired, gin.Error{})
		}
		return nil, response.NewAppError(http.StatusUnauthorized, response.ErrInvalidToken, gin.Error{})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, response.NewAppError(http.StatusUnauthorized, response.ErrCanNotDetectMapClaims, gin.Error{})
	}

	return claims, nil
}
