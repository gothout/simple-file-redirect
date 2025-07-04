package middleware

import (
	"net/http"
	env "simple-file-redirect/internal/configuration/env/application"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		// Extrai o token do header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Compara com a variável de ambiente
		expectedToken := env.GetTokenApp()
		if expectedToken == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server misconfigured: AUTH_TOKEN not set"})
			return
		}

		if token != expectedToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Continua para o handler
		c.Next()
	}
}
