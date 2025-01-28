package main

import (
	"facial-recognition/src/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Set trusted proxies (example: only trust localhost)
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalf("could not set trusted proxies: %s\n", err)
	}

	// * Group v1 routes
	api := r.Group("/api")
	{
		// Group v1 routes
		v1 := api.Group("/v1")
		{
			// Test routes
			test := v1.Group("/test")
			{
				routes.TestRouter(test)
			}
			// Face routes
			face := v1.Group("/face")
			{
				routes.FaceRouter(face)
			}
		}
	}

	// Custom 404 handler
	r.NoRoute(routes.NotFoundRouter)

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
