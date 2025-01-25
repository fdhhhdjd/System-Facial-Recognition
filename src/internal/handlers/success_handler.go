package handlers

import (
	"facial-recognition/src/internal/models"
	pkg "facial-recognition/src/pkg/constants"

	"github.com/gin-gonic/gin"
)

// NewSuccessResponse creates a new SuccessResponse
func CreateSuccessResponse(message string, statusCode int, metadata interface{}) *models.SuccessResponse {
	return &models.SuccessResponse{
		Message:  message,
		Status:   statusCode,
		Metadata: metadata,
	}
}

// Ok sends a 200 OK response
func Ok(c *gin.Context, message string, metadata interface{}) {
	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusOK)
	}
	response := CreateSuccessResponse(message, pkg.StatusOK, metadata)
	Send(c, response.Status, response)
}

// Created sends a 201 Created response
func Created(c *gin.Context, message string, metadata interface{}) {
	if message == "" {
		message = pkg.GetReasonPhrase(pkg.StatusCreated)
	}
	response := CreateSuccessResponse(message, pkg.StatusCreated, metadata)
	Send(c, response.Status, response)
}
