package handlers

import (
	pkg "facial-recognition/src/pkg/constants"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message          string      `json:"message"`
	Status           int         `json:"status"`
	ReasonStatusCode string      `json:"reason_status_code"`
	Option           interface{} `json:"option,omitempty"`
	Metadata         interface{} `json:"metadata,omitempty"`
}

// NewSuccessResponse creates a new SuccessResponse
func CreateSuccessResponse(message string, statusCode int, metadata interface{}) *SuccessResponse {
	return &SuccessResponse{
		Message:  message,
		Status:   statusCode,
		Metadata: metadata,
	}
}

func Ok(c *gin.Context, message string, metadata interface{}) {
	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusOK)
	}
	response := CreateSuccessResponse(message, pkg.StatusOK, metadata)
	Send(c, response.Status, response)
}

func Created(c *gin.Context, message string, metadata interface{}) {
	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusCreated)
	}
	response := CreateSuccessResponse(message, pkg.StatusCreated, metadata)
	Send(c, response.Status, response)
}
