package middleware

import (
	"net/http"
	env "simple-file-redirect/internal/configuration/env/application"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string

		// 1. Verifica se há token na query string
		queryToken := c.Query("token")
		if queryToken != "" {
			token = queryToken
		} else {
			// 2. Caso não tenha, verifica o Authorization header
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header or token query parameter required"})
				return
			}
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}

		// 3. Valida o token
		expectedToken := env.GetTokenApp()
		if expectedToken == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server misconfigured: AUTH_TOKEN not set"})
			return
		}

		if token != expectedToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Next()
	}
}
