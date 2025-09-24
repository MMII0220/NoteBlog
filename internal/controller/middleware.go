package controller

import (
	// "fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	// "myasd/internal/service"
)

// AuthMiddleware проверяет Access Token
func (contr *ControllerStruct) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка
		authHeader := c.GetHeader("Authorization")
		// fmt.Printf("DEBUG: Authorization header: '%s'\n", authHeader)

		if authHeader == "" {
			// fmt.Println("DEBUG: Authorization header is empty")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Достаём сам токен
		tokenParts := strings.Split(authHeader, " ")
		// fmt.Printf("DEBUG: Token parts: %v\n", tokenParts)

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			// fmt.Printf("DEBUG: Invalid token format. Parts count: %d, First part: '%s'\n", len(tokenParts), tokenParts[0])
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		token := tokenParts[1]
		// fmt.Printf("DEBUG: Extracted token: '%s'\n", token)

		// Проверяем токен
		userID, err := contr.serv.ValidateAccessToken(token)
		if err != nil {
			// fmt.Printf("DEBUG: Token validation failed: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// fmt.Printf("DEBUG: Token validated successfully. UserID: %d\n", userID)
		// Кладём userID в контекст
		c.Set("user_id", userID)
		c.Next()
	}
}

func (contr *ControllerStruct) LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		status := c.Writer.Status()
		duration := time.Since(start)

		contr.logger.Info().
			Str("method", method).
			Str("path", path).
			Int("status", status).
			Dur("latency", duration).
			Msg("handled request")
	}
}
