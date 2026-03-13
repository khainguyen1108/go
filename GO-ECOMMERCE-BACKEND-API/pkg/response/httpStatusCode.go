package response

const (
	CodeSuccess         = 20001 //Success
	ErrCodeParamInvalid = 20003 //Email is invalid
	ErrInvalidToken     = 30001 //Token is invalid
)

// message
var msg = map[int]string{
	CodeSuccess:         "success",
	ErrCodeParamInvalid: "param invalid",
	ErrInvalidToken:     "token is invalid",
}
