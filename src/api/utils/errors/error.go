package errors

import "net/http"

// APIError interface
type APIError interface {
	GetStatus() int
	GetMessage() string
	GetError() string
}

type apiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func (a *apiError) GetStatus() int {
	return a.Status
}

func (a *apiError) GetMessage() string {
	return a.Message
}

func (a *apiError) GetError() string {
	return a.Error
}

// NotFound error struct
func NotFound(message string) APIError {
	return &apiError{
		Status:  http.StatusNotFound,
		Message: message,
	}
}

// NewAPIError creates a new API error
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		Status:  statusCode,
		Message: message,
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
