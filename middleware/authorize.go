package middleware

import (
	"fmt"
	"my-blog-api/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")

		// Check if Authorization header is missing
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token format. Use 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		// Extract the token
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := services.JWTAuthService().ValidateToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token validation failed",
			})
			c.Abort()
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			// You can access claims here if needed
			// claims["some_claim_key"]
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid",
			})
			c.Abort()
			return
		}
	}
}
