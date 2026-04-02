package response

const (
	CodeSuccess         = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid

	ErrInvalidToken     = 30001 //Token is invalid
	ErrValidationFailed = 30002

	// User related errors
	ErrUserNotFound            = 40001
	ErrUserAlreadyUnregistered = 40002
	ErrPasswordNotMatch        = 40003
	ErrTokenNotFound           = 40004
	ErrCanNotDetectMapClaims   = 40005
	ErrAlogrithmNotSupported   = 40006
	ErrTokenExpired            = 40007
	ErrInvalidRefreshToken     = 40008
	ErrInvalidSessionId        = 40009
	ErrUserAlreadyBlocked      = 40010

	// Internal errors

	ErrInternalError              = 50001
	ErrGenerateTokenFailed        = 50004
	ErrCreateAccountSessionFailed = 50005
	ErrLogoutFailed               = 50006
	ErrUpdateAccountSessionFailed = 50007
	ErrRevokeAccountFailed        = 50008
)

// message
var msg = map[int]string{
	CodeSuccess:                   "success",
	ErrCodeParamInvalid:           "param invalid",
	ErrInvalidToken:               "token is invalid",
	ErrValidationFailed:           "validation failed",
	ErrInternalError:              "internal server error",
	ErrUserNotFound:               "user not found",
	ErrUserAlreadyUnregistered:    "user already unregistered",
	ErrPasswordNotMatch:           "password not match",
	ErrGenerateTokenFailed:        "failed to generate token",
	ErrCreateAccountSessionFailed: "failed to create account session",
	ErrTokenNotFound:              "token not found",
	ErrCanNotDetectMapClaims:      "cannot detect map claims from token",
	ErrAlogrithmNotSupported:      "signing algorithm is not supported",
	ErrTokenExpired:               "token is expired",
	ErrLogoutFailed:               "failed to logout",
	ErrInvalidRefreshToken:        "invalid refresh token",
	ErrInvalidSessionId:           "invalid session id",
	ErrUpdateAccountSessionFailed: "failed to update account session",
	ErrRevokeAccountFailed:        "failed to revoke account session",
	ErrUserAlreadyBlocked:         "user is already blocked please login again",
}
