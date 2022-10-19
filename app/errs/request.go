package errs

import "net/http"

type RequestError struct {
	StatusCode int
	Message    string
}

func (r RequestError) Error() string {
	return r.Message
}

func NotFound(message string) error {
	return RequestError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func Conflict(message string) error {
	return RequestError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

func NotAcceptable(message string) error {
	return RequestError{
		StatusCode: http.StatusNotAcceptable,
		Message:    message,
	}
}

func UnexpectedError() error {
	return RequestError{
		StatusCode: http.StatusInternalServerError,
		Message:    "unexpected error",
	}
}
