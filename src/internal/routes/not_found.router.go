package routes

import (
	"facial-recognition/src/internal/handlers"
	pkg "facial-recognition/src/pkg/constants"

	"github.com/gin-gonic/gin"
)

func NotFoundRouter(c *gin.Context) {
	handlers.NotFoundError(c, pkg.StatusNotFound)
}
