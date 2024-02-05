package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userClaim, exists := claims["User"]
			if !exists {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token: no User claim"})
				c.Abort()
				return
			}

			userMap, ok := userClaim.(map[string]interface{})
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token: User claim is not a map"})
				c.Abort()
				return
			}

			emailClaim, exists := userMap["Email"]
			if !exists {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token: no Email claim in User"})
				c.Abort()
				return
			}

			email, ok := emailClaim.(string)
			if !ok || email == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token: Email claim is not a string"})
				c.Abort()
				return
			}

			c.Set("userEmail", email)
		}

		c.Next()
	}
}