package errors

import (
	"encoding/json"
	"net/http"
)

// APIError interface
type APIError interface {
	GetStatus() int
	GetMessage() string
	GetErrors() []error
}

type apiError struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Errors  []error `json:"errors,omitempty"`
}

func (a *apiError) GetStatus() int {
	return a.Status
}

func (a *apiError) GetMessage() string {
	return a.Message
}

func (a *apiError) GetErrors() []error {
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
func NewAPIError(statusCode int, message string, errors []error) APIError {
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

// NewAPIErrorFromBytes returns the api error struct for the given byte slice
func NewAPIErrorFromBytes(bytes []byte) (APIError, error) {
	var result apiError
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}
	return &result, nil
}