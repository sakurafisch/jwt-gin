package middleware

import (
	"fmt"
	"net/http"

	"jwt-gin/service.go"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"

		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Failed to execute service.JWTAuthService().ValidateToken(tokenString)",
			})
		}
		if !token.Valid {
			fmt.Println("Invalid token")
			c.JSON(http.StatusOK, gin.H{
				"msg": "Invalid token",
			})
		}
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		return
	}
}
