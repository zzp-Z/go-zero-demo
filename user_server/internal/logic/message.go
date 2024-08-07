package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

// AppError 定义自定义错误类型
type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return e.Code + ": " + e.Message
}
func NewAppError(ctx context.Context, code string, message string, err error) *AppError {
	logx.WithContext(ctx).Errorf(code, "：", message, " err：", err)
	return &AppError{Code: code, Message: message}
}
