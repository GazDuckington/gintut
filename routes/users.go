package routes

import (
	"gintut/helpers/authenticator"
	"gintut/helpers/controllers"
	"gintut/models"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
)

func ping(c *gin.Context) {
	claims := authenticator.ClaimsHandler(c)
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
