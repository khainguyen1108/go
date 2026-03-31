package response

const (
	CodeSuccess         = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid
	ErrInvalidToken     = 30001 //Token is invalid
	ErrValidationFailed = 30002
	ErrInternalError    = 50001

	// User related errors
	ErrUserNotFound = 40001
)

// message
var msg = map[int]string{
	CodeSuccess:         "success",
	ErrCodeParamInvalid: "param invalid",
	ErrInvalidToken:     "token is invalid",
	ErrValidationFailed: "validation failed",
	ErrInternalError:    "internal server error",
	ErrUserNotFound:     "user not found",
}
