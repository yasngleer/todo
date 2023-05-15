package api

import (
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, my_user_id string) {
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte("A String Very Very Very Strong!!@##$!@#$"))
			return b, nil
		})
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user_id := claims["id"].(string)
			idint, _ := strconv.Atoi(user_id)
			c.Set("user_id", idint)
		}
	}
}
