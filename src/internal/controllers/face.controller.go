package controllers

import (
	"facial-recognition/src/internal/handlers"
	"facial-recognition/src/internal/services"

	"github.com/gin-gonic/gin"
)

func DetectFaces(c *gin.Context) error {
	result := services.DetectFaces(c)
	if result == nil {
		return nil
	}

	handlers.Ok(c, "Detect Faces", result)
	return nil
}

func RegisterFace(c *gin.Context) error {
	result := services.RegisterFace(c)
	if result == nil {
		return nil
	}

	handlers.Created(c, "Register Face", result)
	return nil
}
