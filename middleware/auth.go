package middleware

import (
	"fmt"
	"lms/database"
	"lms/models"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("User")
		if exists && user.(models.CrmUsers).Id != 0 {
			return
		} else {
			log.Error().Str("RequestURI", c.Request.RequestURI).Msg("JWT auth errors")

			err, _ := c.Get("authErr")
			if err != nil {
				_ = c.AbortWithError(http.StatusUnauthorized, err.(error))
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Auth error"})
			}
		}

	}
}

func JwtTokenVerify() gin.HandlerFunc {

	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		if bearer != "" {
			jwtString := strings.Split(bearer, " ")
			if len(jwtString) == 2 {
				jwtEncoded := jwtString[1]
				token, err := jwt.Parse(jwtEncoded, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signin method %v", token.Header["alg"])
					}
					secret := []byte(os.Getenv("JWT_SECRET"))
					return secret, nil
				})

				if err != nil {
					log.Debug().Msg(err.Error())
					return
				}
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					userId := uint(claims["user_id"].(float64))

					var user models.CrmUsers
					if userId != 0 {
						database := database.GetConnection()
						database.First(&user, userId)
					}

					c.Set("User", user)
					c.Set("userId", user.Id)
				}
			}
		}
	}
}
