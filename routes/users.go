package routes

import (
	// "fmt"
	"gintut/helpers/authenticator"
	"gintut/helpers/controllers"
	"gintut/models"
	"net/http"
	// "strings"

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

func GetUserByEmail(c *gin.Context) {
	// var personal models.TPersonal
	// Get the email from the URL parameter
	email := c.Param("email")

	// Retrieve personal details by email
	// eml_field := fmt.Sprintf(`"%s"`, strings.ToUpper("EML"))
	person, err := authenticator.GetPersonalDetail("eml", email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})
}

// get all business unit
func GetAllBusinessUnits(c *gin.Context) {
	var businessUnit []models.TBusinessunit
	controllers.GetTable(c, businessUnit)
}
