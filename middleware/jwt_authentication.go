package middleware

import (
	"Glossika_interview/config"
	"Glossika_interview/database/models"
	"Glossika_interview/myerror"
	"Glossika_interview/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"strings"
)

var jwtKey = []byte(config.GetString("jwt.secret"))

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.Error(c, myerror.NewTokenError(myerror.TYPE_HEADER_MISSING))
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerSchema)
		if tokenString == authHeader {
			response.Error(c, myerror.NewTokenError(myerror.TYPE_TOKEN_INVALID))
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			response.Error(c, myerror.NewTokenError(myerror.TYPE_TOKEN_INVALID))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			db := c.MustGet("db").(*gorm.DB)
			var user models.User
			db.First(&user, int(claims["user_id"].(float64)))
			c.Set("user", user)
		}

		// Token is valid
		c.Next()
	}
}
