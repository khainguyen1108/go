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
