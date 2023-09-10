package error

import "net/http"

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Cause   error  `json:"cause"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewNotFoundError(message string, cause error) *ApiError {
	if message == "" {
		message = "Api not found"
	}

	return &ApiError{
		Status:  http.StatusNotFound,
		Message: message,
		Cause:   cause,
	}
}

func NewBadRequestError(message string, cause error) *ApiError {
	if message == "" {
		message = "Something went wrong while processing your request."
	}

	return &ApiError{
		Status:  http.StatusBadRequest,
		Message: message,
		Cause:   cause,
	}
}

func NewUnauthorizedError(message string, cause error) *ApiError {
	if message == "" {
		message = "Missing or invalid authentication."
	}

	return &ApiError{
		Status:  http.StatusUnauthorized,
		Message: message,
		Cause:   cause,
	}
}

func NewForbiddenError(message string, cause error) *ApiError {
	if message == "" {
		message = "You are not allowed to perform this request."
	}

	return &ApiError{
		Status:  http.StatusForbidden,
		Message: message,
		Cause:   cause,
	}
}
