package controllers

import (
	"gintut/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAll is a generic function to fetch all records of a given model.
func GetTable(c *gin.Context, model interface{}) {

	// Get the DB instance from Gin's context
	db := c.MustGet("db").(*gorm.DB)

	// Dynamic filter
	// Assuming helpers.DynamicFilter accepts a *gorm.DB and returns a *gorm.DB
	db = helpers.DynamicFilter(db, c)

	// Pagination
	// Assuming helpers.Paginate returns a paginationResult with a *gorm.DB and Pagination data
	paginationResult := helpers.Paginate(db, c, model)

	// Retrieve all records from the table
	if err := paginationResult.DB.Find(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Access pagination data
	pagination := paginationResult.Pagination

	c.JSON(http.StatusOK, gin.H{
		"data":       model,
		"pagination": pagination,
	})
}
