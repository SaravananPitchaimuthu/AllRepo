package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
