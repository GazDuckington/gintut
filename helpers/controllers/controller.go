package controllers

import (
	"fmt"
	"gintut/helpers"
	"gintut/initializers"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// query a model
func GetModel(model interface{}, fieldName string, fieldValue interface{}) error {
	initializers.InitEnv()
	db := initializers.InitDb()

	// Create a new instance of the model
	instance := reflect.New(reflect.TypeOf(model).Elem()).Interface()

	// Retrieve the single record
	if err := db.Where(fieldName+" = ?", fieldValue).First(instance).Error; err != nil {
		return err
	}

	// Copy the data from the instance to the provided model pointer
	reflect.ValueOf(model).Elem().Set(reflect.ValueOf(instance).Elem())

	return nil
}

// generic function to fetch all records of a given model.
func GetTable(c *gin.Context, model interface{}) {
	var order string

	// Get the DB instance from Gin's context
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
