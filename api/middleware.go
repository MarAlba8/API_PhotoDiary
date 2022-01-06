package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtKey = []byte("my_secret_key")
		const prefix = "Bearer "

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(400, gin.H{"error": "Authorization header missing or empty"})
			return
		}

		token := header[len(prefix):]
		tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.IndentedJSON(
					http.StatusUnauthorized,
					gin.H{"message": "Signature invalid"})
				return
			}
			c.AbortWithStatusJSON(401, gin.H{"error": fmt.Sprintf("unable to verify token with error: %v", err)})
		}

		if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
			fmt.Println(claims["identifier"])
			fmt.Println(claims)
			c.Set("identifier", claims["identifier"])
		} else {
			fmt.Println(err)
		}

		c.Next()
	}
}
