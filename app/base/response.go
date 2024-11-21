package base

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BaseApiController struct{}

// SendResponse formats a successful response
func (b *BaseApiController) SendResponse(c *gin.Context, result interface{}, message ...string) {
	response := gin.H{
		"success": true,
		"data":    result,
	}

	// Check if there is a message, add it to the response
	if len(message) > 0 && message[0] != "" {
		response["message"] = message[0]
	}

	c.JSON(http.StatusOK, response)
}

// SendError formats an error response
func (b *BaseApiController) SendError(c *gin.Context, message string, code int) {
	if code == 0 {
		code = http.StatusNotFound
	}

	response := gin.H{
		"success": false,
		"message": message,
	}

	c.JSON(code, response)
}

// SendValidator formats validation error response
func (b *BaseApiController) SendValidator(c *gin.Context, errors map[string]string) {
	response := gin.H{
		"success": false,
		"errors": errors,
	}

	c.JSON(http.StatusUnprocessableEntity, response)
}

// SendPaginationResponse formats a paginated response
func (b *BaseApiController) SendPaginationResponse(c *gin.Context, result interface{}, pagination map[string]interface{}, message string) {
	response := gin.H{
		"success":    true,
		"data":       result,
		"pagination": pagination,
	}

	if message != "" {
		response["message"] = message
	}

	c.JSON(http.StatusOK, response)
}

func (b *BaseApiController) SendPaginationArrayResponse(c *gin.Context, items []interface{}, total int64, pageNum int, perPageNum int, message string) {
    pageNum64 := int64(pageNum)
    perPageNum64 := int64(perPageNum)

    pagination := gin.H{
        "total":        total,
        "per_page":     perPageNum64,
        "current_page": pageNum64,
        "last_page":    (total + perPageNum64 - 1) / perPageNum64,
    }

    b.SendPaginationResponse(c, items, pagination, message)
}

func (b *BaseApiController) ParsePaginationParams(c *gin.Context) (int, int) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

    if page <= 0 {
        page = 1
    }

    if perPage <= 0 {
        perPage = 10
    }

    return page, perPage
}