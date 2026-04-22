package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ok := c.GetHeader("Authorization")
		if ok == "" {
			c.AbortWithStatus(401)
			return
		}
		res := strings.TrimPrefix(ok, "Bearer ")

		token, err := jwt.Parse(res, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil

		})
		_ = token
		if err != nil {
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
