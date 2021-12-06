package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int 			`json:"code,omitempty"`
	StatusText string 		`json:"status,omitempty"`
	Message string 			`json:"error,omitempty"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 400,
		StatusText: "Bad request",
		Message: err.Error(),
	}
}

func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: 500,
		StatusText: "Internal error",
		Message: err.Error(),
	}
}