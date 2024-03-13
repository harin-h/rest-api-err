package restapierr

import "net/http"

type AppError struct {
	Code     int
	Message  string
	LogError string
}

func InternalServerError(logError string) AppError {
	return AppError{Code: http.StatusInternalServerError, Message: "unexpected error", LogError: logError}
}
