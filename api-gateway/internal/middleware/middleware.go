package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Middleware для аутентификации с JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			c.Abort()
			return
		}

		token := authHeader[len("Bearer "):]
		// Здесь ты можешь добавить проверку токена через твою систему авторизации
		if token != "your-valid-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
