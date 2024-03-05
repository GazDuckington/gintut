package routes

import (
	"gintut/helpers/authenticator"
	"gintut/helpers/controllers"
	"gintut/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Email string `json:"email" form:"email"`
	Pswd  string `json:"password" form:"password"`
}

func Login(c *gin.Context) {
	var personal models.TPersonal
	var user models.TUsers
	var auth Auth

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controllers.GetModel(&user, `"EML"`, auth.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controllers.GetModel(&personal, `"NIK"`, user.Nik)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if authenticator.CheckPassword(user.Pswd, auth.Pswd) {
		token, err := authenticator.GenerateJWT(user.Nik)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "personal_data": personal})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login failed"})
	}

}
