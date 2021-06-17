package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/users-api/domain"
	"strconv"
)

func GenerateValidation(c *gin.Context) domain.Pagination {
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	// set default value to limit if limit zero
	if limit == 0 {
		limit = 10
	}
	return domain.Pagination{Limit: limit, Offset:  offset}
}