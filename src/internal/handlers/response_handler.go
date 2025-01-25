package handlers

import (
	"github.com/gin-gonic/gin"
)

func Send(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
