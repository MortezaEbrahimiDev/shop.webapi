package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"shop.go/internal/auth"
	"strings"
)

//import "github.com/gin-gonic/gin"

func Authorize(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := &auth.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey(), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// بررسی نقش

		for _, role := range allowedRoles {
			for _, allowed := range allowedRoles {
				if role == allowed {
					c.Set("userId", claims.UserId)
					c.Next()
					return
				}
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
	}
}
