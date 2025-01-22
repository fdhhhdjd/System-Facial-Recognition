package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AsyncHandler wraps a handler function and handles errors asynchronously
func AsyncHandler(fn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Run the handler function and capture any error
		if err := fn(c); err != nil {
			// Log the error with request details
			log.Printf("Error: %v | Method: %s | Path: %s | IP: %s", err, c.Request.Method, c.Request.URL.Path, c.ClientIP())

			// Respond with an error message
			InternalServerError(c, http.StatusInternalServerError, err.Error())

			// Abort the request to prevent further processing
			c.Abort()
			return
		}

		// Continue to the next middleware/handler
		c.Next()
	}
}
