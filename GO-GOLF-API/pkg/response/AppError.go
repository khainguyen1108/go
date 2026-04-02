package response

type AppError struct {
	StatusCode int
	AppCode    int
	Message    string
	RootErr    error
}

func (e *AppError) Error() string {
	if e.RootErr != nil {
		return e.RootErr.Error()
	}
	return e.Message
}

func NewAppError(status int, appCode int, err error) *AppError {
	return &AppError{
		StatusCode: status,
		AppCode:    appCode,
		Message:    "",
		RootErr:    err,
	}
}

func NewAppErrorWithMessage(status int, appCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: status,
		AppCode:    appCode,
		Message:    message,
		RootErr:    err,
	}
}
