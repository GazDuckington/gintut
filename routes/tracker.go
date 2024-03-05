package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetServerTime(c *gin.Context) {
	serverTime := time.Now().Format("Mon Jan 02 2006 15:04:05 -0700")
	c.JSON(http.StatusOK, gin.H{"server_time": serverTime})
}
