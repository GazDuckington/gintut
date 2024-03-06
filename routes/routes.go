package routes

import (
	"github.com/gin-gonic/gin"
)

// route group to be protected by bearer token
func RouteGroupProtected(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	router.GET("/ping", ping)
	userGroup.GET("/list", GetAllPersonal)
	userGroup.GET("/detail/:email", GetUserByEmail)
	router.GET("/bu", GetAllBusinessUnits)
}

// route group to be unprotected by bearer token
func RouteGroupUnprotected(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	authGroup.POST("/login", Login)
	router.GET("/get_server_time", GetServerTime)
}
