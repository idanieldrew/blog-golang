package restError

import "net/http"

type RestError struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NotFoundError Not found error
func NotFoundError(message string) *RestError {
	return &RestError{
		Error:   "not_found",
		Status:  http.StatusNotFound,
		Message: message,
	}
}

// ServerError Server error
func ServerError(message string) *RestError {
	return &RestError{
		Error:   "server_error",
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

// ValidationError Validation error
func ValidationError(message string) *RestError {
	return &RestError{
		Error:   "validation_error",
		Status:  http.StatusUnprocessableEntity,
		Message: message,
	}
}

// BadRequestError Bad request error
func BadRequestError(message string) *RestError {
	return &RestError{
		Error:   "bad_request",
		Status:  http.StatusBadRequest,
		Message: message,
	}
}

// UnauthorizedError Unauthorized
func UnauthorizedError(message string) *RestError {
	return &RestError{
		Error:   "unauthorized",
		Status:  http.StatusUnauthorized,
		Message: message,
	}
}
