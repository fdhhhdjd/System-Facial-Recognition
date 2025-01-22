package handlers

import (
	"time"

	pkg "facial-recognition/src/pkg/constants"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Now     int64  `json:"now"`
}

func CreateErrorResponse(message string, status int, code int) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Status:  status,
		Now:     time.Now().Unix(),
	}
}

// BadRequestError represents a 400 Bad Request error
func BadRequestError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusBadRequest)
	}
	response := CreateErrorResponse(message, pkg.StatusBadRequest, code)
	Send(c, response.Status, response)
}

// NotFoundError represents a 404 Not Found error
func NotFoundError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusNotFound)
	}
	response := CreateErrorResponse(message, pkg.StatusNotFound, code)
	Send(c, response.Status, response)
}

func InternalServerError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusInternalServerError)
	}
	response := CreateErrorResponse(message, pkg.StatusInternalServerError, code)
	Send(c, response.Status, response)
}
