package restError

import "net/http"

type RestError struct {
	Error   string `json:"error"`
	Status  uint   `json:"status"`
	Message string `json:"message"`
}

func NotFoundError(message string) *RestError {
	return &RestError{
		Error:   "not_found",
		Status:  http.StatusNotFound,
		Message: message,
	}
}

func ServerError(message string) *RestError {
	return &RestError{
		Error:   "server_error",
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}
