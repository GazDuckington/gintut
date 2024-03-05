package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginationResult struct {
	DB         *gorm.DB
	Pagination map[string]interface{}
	Count      int
}

func Paginate(db *gorm.DB, c *gin.Context, m interface{}) *PaginationResult {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	offset := (page - 1) * pageSize

	var count int64
	db.Model(m).Count(&count)

	pagination := map[string]interface{}{
		"page":      page,
		"page_size": pageSize,
		"count":     int(count),
	}

	return &PaginationResult{
		DB:         db.Offset(offset).Limit(pageSize),
		Pagination: pagination,
	}
}

func DynamicFilter(db *gorm.DB, c *gin.Context) *gorm.DB {
	qParams := c.Request.URL.Query()

	// Remove pagination parameters
	qParams.Del("page")
	qParams.Del("page_size")
	qParams.Del("order_by")

	// Create dynamic filters
	var dynamicFilters []string
	var queryParams []interface{}

	for param, value := range qParams {
		words := strings.Fields(value[0])
		var filterStrings []string
		for _, word := range words {
			filterStrings = append(filterStrings, fmt.Sprintf(`"%s" = ?`, strings.ToUpper(param)))
			queryParams = append(queryParams, word)
		}
		dynamicFilters = append(dynamicFilters, strings.Join(filterStrings, " OR "))
	}

	// Construct the final dynamic filter query
	dynamicQuery := strings.Join(dynamicFilters, " AND ")

	// Apply dynamic filter to the query
	return db.Where(dynamicQuery, queryParams...)

}
