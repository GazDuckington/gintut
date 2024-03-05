package authenticator

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the value of the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Check if the Authorization header starts with "Bearer"
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Extract the token from the header
		authToken := authHeaderParts[1]

		// Validate the token (you should implement your own token validation logic here)
		if !isValidToken(authToken) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// If the token is valid, proceed to the next middleware or handler
		c.Next()
	}
}

func isValidToken(tokenString string) bool {
	var secret = os.Getenv("SUPER_SECRET")
	if secret == "" {
		// Handle the case when SUPER_SECRET is not set
		log.Println("SUPER_SECRET environment variable is not set")
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		// Handle the case when there's an error parsing the token
		log.Println("Error parsing token:", err)
		return false
	}

	return token.Valid
}
