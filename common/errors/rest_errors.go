package errors

import (
	"fmt"
	"net/http"
)

type RestError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restError struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	err     string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e *restError) Message() string {
	return e.message
}

func (e *restError) Status() int {
	return e.status
}

func (e *restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", e.message, e.status, e.err, e.causes)
}

func (e *restError) Causes() []interface{} {
	return e.causes
}

func NewRestError(message string, status int, errorMsg string, causes []interface{}) RestError {
	return &restError{
		message: message,
		status:  status,
		err:     errorMsg,
		causes:  causes,
	}
}

func NewBadRequestError(message string) RestError {
	return &restError{
		message: message,
		status:  http.StatusBadRequest,
		err:     "bad_request",
	}
}

func NewUnauthorizedError(message string) RestError {
	return &restError{
		message: message,
		status:  http.StatusUnauthorized,
		err:     "unauthorized",
	}
}

func NewNotFoundError(message string) RestError {
	return &restError{
		message: message,
		status:  http.StatusNotFound,
		err:     "not_found",
	}
}

func NewInternalServerError(message string, err error) RestError {
	result := &restError{
		message: message,
		status:  http.StatusInternalServerError,
		err:     "internal_server_error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
