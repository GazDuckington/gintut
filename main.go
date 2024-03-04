package main

import (
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
		log.Println("Middleware: Setting db in context")
		c.Set("db", db)
		c.Next()
	})

	routes.InitializeRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080
}
