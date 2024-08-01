package logic

// AppError 定义自定义错误类型
type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Code + ": " + e.Message
}
func NewAppError(code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}
