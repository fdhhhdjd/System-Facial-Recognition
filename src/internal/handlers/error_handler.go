package handlers

import (
	"time"

	"facial-recognition/src/internal/models"
	pkg "facial-recognition/src/pkg/constants"

	"github.com/gin-gonic/gin"
)

// Send sends a response
func CreateErrorResponse(message string, status int, code int) *models.ErrorResponse {
	return &models.ErrorResponse{
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

// TooManyRequestsError represents a 429 Too Many Requests error
func TooManyRequestsError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusTooManyRequests)
	}
	response := CreateErrorResponse(message, pkg.StatusTooManyRequests, code)
	Send(c, response.Status, response)
}

// ForbiddenError represents a 403 Forbidden error
func ForbiddenError(c *gin.Context, code int, messages ...string) {
	message := ""
	if len(messages) > 0 {
		message = messages[0]
	}

	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusForbidden)
	}
	response := CreateErrorResponse(message, pkg.StatusForbidden, code)
	Send(c, response.Status, response)
}

// InternalServerError represents a 500 Internal Server Error
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
