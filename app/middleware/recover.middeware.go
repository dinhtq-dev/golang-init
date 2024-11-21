package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error for debugging purposes
				log.Printf("Panic recovered: %v", err)

				// Set the response header
				c.Header("Content-Type", "application/json")

				// Return a JSON response with an appropriate status code
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error":   "Something went wrong",
					"message": err,
					"status":  false,
				})
			}
		}()

		c.Next()
	}
}
