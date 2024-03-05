package routes

import (
	"github.com/gin-gonic/gin"
)

// RouteGroupProtected sets up the album routes
func RouteGroupProtected(router *gin.RouterGroup) {
	router.GET("/ping", ping)
	router.GET("/personal", GetAllPersonal)
}

func RouteGroupUnprotected(router *gin.RouterGroup) {
	router.POST("/login", Login)
}
