package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{Message: e.Message, Code: e.Code}
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
