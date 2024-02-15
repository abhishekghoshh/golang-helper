package error

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
func New(message string, statusCode int) *AppError {
	return &AppError{
		Message: message,
		Code:    statusCode,
	}
}
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}
func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}

func NewFunctionNotImplementedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotImplemented,
	}
}

func (appError AppError) Error() string {
	return appError.Message
}
