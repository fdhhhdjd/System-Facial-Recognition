package routes

import (
	"github.com/gin-gonic/gin"
)

func TestRouter(router *gin.RouterGroup) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
