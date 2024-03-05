package routes

import (
	"gintut/helpers/authenticator"
	"gintut/helpers/controllers"
	"gintut/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// test only
func ping(c *gin.Context) {
	claims := authenticator.GetClaimsData(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"claims":  claims,
	})
}

// get all personal data
func GetAllPersonal(c *gin.Context) {
	var personal []models.TPersonal
	controllers.GetTable(c, personal)
}

// get all business unit
func GetAllBusinessUnits(c *gin.Context) {
	var businessUnit []models.TBusinessunit
	controllers.GetTable(c, businessUnit)
}
