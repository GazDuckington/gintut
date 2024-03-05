package routes

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes sets up the album routes
func InitializeRoutes(router *gin.Engine) {
	router.GET("/ping", ping)
	router.GET("/personal", GetAllPersonal)
	router.POST("/login", Login)
}
