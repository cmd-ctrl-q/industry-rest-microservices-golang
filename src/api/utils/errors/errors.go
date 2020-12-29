package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	SomeStatus  int    `json:"status"`
	SomeMessage string `json:"message"`
	SomeError   string `json:"error,omitempty"` // if no error, dont show it in final json
}

func (e *apiError) Status() int {
	return e.SomeStatus
}

func (e *apiError) Message() string {
	return e.SomeMessage
}

func (e *apiError) Error() string {
	return e.SomeError
}

func NewApiError(statusCode int, message string) APIError {
	return &apiError{
		SomeStatus:  statusCode,
		SomeMessage: message,
	}
}

func NewApiErrorFromBytes(body []byte) (APIError, error) {
	var result apiError
	// if we have any error trying to use this []byte array when populating
	// it into the result, then throw error
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

func NewInternalServerError(message string) APIError {
	return &apiError{
		SomeStatus:  http.StatusInternalServerError,
		SomeMessage: message,
	}
}

func NewNotFoundError(message string) APIError {
	return &apiError{
		SomeStatus:  http.StatusNotFound,
		SomeMessage: message,
	}
}

func NewBadRequestError(message string) APIError {
	return &apiError{
		SomeStatus:  http.StatusBadRequest,
		SomeMessage: message,
	}
}
