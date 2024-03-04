package routes

import (
	"fmt"
	"gintut/helpers"
	"gintut/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetAllPersonal(c *gin.Context) {
	var personal []models.TPersonal

	// Get the DB instance from Gin's context
	db := c.MustGet("db").(*gorm.DB)
	db = db.Where(fmt.Sprintf(`"ENDDA" > ?`), helpers.Today)

	// Dynamic filter
	dynamicFilter := helpers.DynamicFilter(db, c)

	// Retrieve all records from the t_personal table
	paginationResult := helpers.Paginate(dynamicFilter, c, &models.TPersonal{})

	// Retrieve all records from the t_personal table
	if err := paginationResult.DB.Find(&personal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Access pagination data
	pagination := paginationResult.Pagination

	c.JSON(http.StatusOK, gin.H{
		"data":       personal,
		"pagination": pagination,
	})
}
