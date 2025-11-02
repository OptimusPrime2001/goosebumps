package utils

import (
	"net/http"
)

type ErrorCode string
type AppError struct {
	Err     error
	Message string
	Code    ErrorCode
}

const (
	ErrCodeBadRequest ErrorCode = "BAD_REQUEST"
	ErrCodeNotFound   ErrorCode = "NOT_FOUND"
	ErrCodeConflict   ErrorCode = "CONFLICT"
	ErrCodeInternal   ErrorCode = "INTERNAL_SERVER_ERROR"
)

func (e *AppError) Error() string {
	return e.Message
}

func NewError(message string, code ErrorCode) error {
	return &AppError{
		Message: message,
		Code:    code,
	}
}
func WrapAppError(systemErr error, message string, code ErrorCode) *AppError {
	return &AppError{
		Err:     systemErr,
		Message: message,
		Code:    code,
	}
}

func httpStatusFromErrorCode(code ErrorCode) int {
	switch code {
	case ErrCodeBadRequest:
		return http.StatusBadRequest
	case ErrCodeNotFound:
		return http.StatusNotFound
	case ErrCodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
