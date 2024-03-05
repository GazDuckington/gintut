package main

import (
	"gintut/helpers/authenticator"
	"gintut/initializers"
	"gintut/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize some stuffs
	initializers.InitEnv()
	db := initializers.InitDb()

	// initialize routes
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		log.Println("Middleware: Setting db in context", db)
		c.Set("db", db)
		c.Next()
	})

	api := router.Group("/api/v1")
	{
		// Apply authentication middleware to RouteGroupProtected routes
		protected := api.Group("/protected")
		protected.Use(authenticator.AuthMiddleware())
		routes.RouteGroupProtected(protected)

		// Include RouteGroupUnprotected routes directly under /api/v1
		routes.RouteGroupUnprotected(api)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}
