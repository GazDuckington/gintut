package routes

import (
	"github.com/gin-gonic/gin"
)

// route group to be protected by bearer token
func RouteGroupProtected(router *gin.RouterGroup) {
	router.GET("/ping", ping)
	router.GET("/personal", GetAllPersonal)
	router.GET("/bu", GetAllBusinessUnits)
}

// route group to be unprotected by bearer token
func RouteGroupUnprotected(router *gin.RouterGroup) {
	router.POST("/login", Login)
	router.GET("/get_server_time", GetServerTime)
}
