package routes

import (
	"facial-recognition/src/internal/controllers"
	"facial-recognition/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

func FaceRouter(router *gin.RouterGroup) {
	// Face detection
	router.POST("/detect", handlers.AsyncHandler(controllers.DetectFaces))

	// Face registration
	router.POST("/register", handlers.AsyncHandler(controllers.RegisterFace))
}
