package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware là một middleware kiểm tra người dùng đã đăng nhập hay chưa
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Giả sử bạn kiểm tra token trong header (bạn có thể thay thế bằng cách kiểm tra session, cookie, etc.)
		token := c.GetHeader("Authorization")
		if token == "" {
			// Nếu không có token, trả về lỗi 401 Unauthorized
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort() // Dừng tiếp tục xử lý request
			return
		}

		// Nếu có token, có thể xác thực nó ở đây (ví dụ: kiểm tra với một dịch vụ xác thực)
		// Nếu token không hợp lệ, trả về lỗi 401 Unauthorized
		if token != "valid_token_example" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			c.Abort() // Dừng tiếp tục xử lý request
			return
		}

		// Nếu token hợp lệ, tiếp tục xử lý request
		c.Next()
	}
}

// LoggingMiddleware là một middleware dùng để ghi log thông tin của mỗi request
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy thông tin request
		method := c.Request.Method
		path := c.Request.URL.Path

		// Ghi log thông tin request
		fmt.Printf("Request %s %s\n", method, path)

		// Tiến hành xử lý request tiếp theo
		c.Next()

		// Sau khi request xử lý xong, bạn có thể lấy status code trả về
		statusCode := c.Writer.Status()
		fmt.Printf("Response %s %s Status Code: %d\n", method, path, statusCode)
	}
}
