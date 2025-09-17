package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"myasd/internal/service"
)

// AuthMiddleware проверяет Access Token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Достаём сам токен
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		token := tokenParts[1]

		// Проверяем токен
		userID, err := service.ValidateAccessToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Кладём userID в контекст
		c.Set("user_id", userID)
		c.Next()
	}
}
