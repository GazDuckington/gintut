package main

import (
	"gintut/helpers/authenticator"
	"gintut/initializers"
	"gintut/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize some stuffs
	initializers.InitEnv()
	db := initializers.InitDb()

	// Set the time zone globally
	loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
	if err != nil {
		// Handle error if the time zone cannot be loaded
		panic(err)
	}
	time.Local = loc

	// initialize routes
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		log.Println("Middleware: Setting db in context", db)
		c.Set("db", db)
		c.Next()
	})

	apiV1 := router.Group("/api/v1")
	{
		// Apply authentication middleware to RouteGroupProtected routes
		protected := apiV1.Group("/protected")
		protected.Use(authenticator.AuthMiddleware())
		routes.RouteGroupProtected(protected)

		// Include RouteGroupUnprotected routes directly under /api/v1
		routes.RouteGroupUnprotected(apiV1)
	}
	router.Run() // listen and serve on 0.0.0.0:8080
}
