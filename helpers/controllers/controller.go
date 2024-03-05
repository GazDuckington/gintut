package controllers

import (
	"fmt"
	"gintut/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAll is a generic function to fetch all records of a given model.
func GetTable(c *gin.Context, model interface{}) {

	// Get the DB instance from Gin's context
	var order string
	orderBy := c.Query("order_by") // Assuming 'order_by' is the query parameter for ordering
	db := c.MustGet("db").(*gorm.DB)

	// Apply ordering if specified
	if orderBy != "" {
		order = fmt.Sprintf(`"%s"`, strings.ToUpper(orderBy))
		db = db.Order(order)
	}
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
