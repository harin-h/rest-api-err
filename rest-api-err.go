package restapierr

import "net/http"

type AppError struct {
	Code     int
	Message  string
	LogError string
}

func (err AppError) Error() string {
	return err.Message
}

func InternalServerError(logError string) error {
	return AppError{Code: http.StatusInternalServerError, Message: "unexpected error", LogError: logError}
}
