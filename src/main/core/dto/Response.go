package dto

import "net/http"

type Response struct {
	Success    bool          `json:"success"`
	Data       interface{}   `json:"data,omitempty"`
	HTTPStatus int           `json:"http_status"`
	Error      *ErrorDetails `json:"error,omitempty"`
}

type ErrorDetails struct {
	Message string `json:"message"`
}

func NewSuccessResponse(data interface{}, httpStatus int) Response {
	return Response{
		Success:    true,
		Data:       data,
		HTTPStatus: httpStatus,
	}
}

func NewErrorResponse(message string, httpStatus int) Response {
	return Response{
		Success:    false,
		HTTPStatus: httpStatus,
		Error: &ErrorDetails{
			Message: message,
		},
	}
}

func NewSuccessResponseOk(data interface{}) Response {
	return NewSuccessResponse(data, http.StatusOK)
}

func NewSuccessResponseCreated(data interface{}) Response {
	return NewSuccessResponse(data, http.StatusOK)
}

func NewErrorResponseBadRequest(message string) Response {
	return NewErrorResponse(message, http.StatusBadRequest)
}

func NewErrorResponseNotFound(message string) Response {
	return NewErrorResponse(message, http.StatusNotFound)
}

func NewErrorResponseInternalServerError(message string) Response {
	return NewErrorResponse(message, http.StatusInternalServerError)
}
