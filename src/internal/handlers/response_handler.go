package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message          string      `json:"message"`
	Status           int         `json:"status"`
	ReasonStatusCode string      `json:"reason_status_code"`
	Option           interface{} `json:"option,omitempty"`
	Metadata         interface{} `json:"metadata,omitempty"`
}

func Send(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}
