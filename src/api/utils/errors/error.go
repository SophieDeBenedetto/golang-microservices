package errors

import (
	"net/http"

	"github.com/SophieDeBenedetto/golang-microservices/src/api/domain/github"
)

// APIError interface
type APIError interface {
	GetStatus() int
	GetMessage() string
	GetErrors() []github.Error
}

type apiError struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Errors  []github.Error `json:"errors,omitempty"`
}

func (a *apiError) GetStatus() int {
	return a.Status
}

func (a *apiError) GetMessage() string {
	return a.Message
}

func (a *apiError) GetErrors() []github.Error {
	return a.Errors
}

// NotFound error struct
func NotFound(message string) APIError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

// NewAPIError creates a new API error
func NewAPIError(statusCode int, message string, errors []github.Error) APIError {
	return &apiError{
		Status:  statusCode,
		Message: message,
		Errors:  errors,
	}
}

// InternalServerError error struct
func InternalServerError(message string) APIError {
	return &apiError{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

// BadRequestError error struct
func BadRequestError(message string) APIError {
	return &apiError{
		Status:  http.StatusBadRequest,
		Message: message,
	}
}
