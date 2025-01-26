package handlers

import (
	"github.com/gin-gonic/gin"
)

// Send a response to the client to json
func Send(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
