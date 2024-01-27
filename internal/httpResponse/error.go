package httpResponse

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message    string  `json:"message"`
	StatusCode int     `json:"status"`
	Error      []Cause `json:"error,omitempty"`
	Instance   string  `json:"instance"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewBadRequestError(message string, r *http.Request) *Error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Instance:   r.URL.Path,
	}
}

func NewBadRequestValidationError(message string, cause []Cause, r *http.Request) *Error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      cause,
		Instance:   r.URL.Path,
	}
}

func NewUnauthorizedError(message string, r *http.Request) *Error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
		Instance:   r.URL.Path,
	}
}

func NewNotFoundError(message string, r *http.Request) *Error {
	return &Error{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Instance:   r.URL.Path,
	}
}

func (e *Error) RenderJSON(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(e.StatusCode)
	json.NewEncoder(w).Encode(e)
}
