package routes

import (
	"gintut/helpers/controllers"
	"gintut/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetAllPersonal(c *gin.Context) {
	var personal []models.TPersonal

	controllers.GetTable(c, personal)
}
