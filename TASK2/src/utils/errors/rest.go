package errors

import "net/http"

type BaseErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *BaseErr {
	return &BaseErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewInternalServerError() *BaseErr {
	return &BaseErr{
		Message: "Something Went Wrong",
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
	}
}

func NewNotFoundError(message string) *BaseErr {
	return &BaseErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found Error",
	}
}
