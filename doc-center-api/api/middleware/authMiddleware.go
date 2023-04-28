package middleware

import (
	// "doc-center-api/api/auth"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"doc-center-api/api/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		result := utils.TokenIsValid(c)
		if !result {
			fmt.Println("token invalid")
			c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "Authentication required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
