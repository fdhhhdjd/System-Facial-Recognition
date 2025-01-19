package routes

import (
	"facial-recognition/src/internal/controllers"
	"facial-recognition/src/internal/handlers"

	"github.com/gin-gonic/gin"
)

func FaceRouter(router *gin.RouterGroup) {
	// Face detection
	router.GET("/detect", handlers.AsyncHandler(controllers.DetectFaces))
}
