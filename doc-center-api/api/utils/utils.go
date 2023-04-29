package utils

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.StandardClaims
	UserEmail string
}

var TOKEN_EXP = time.Now().Add(time.Hour * 3).Unix() // 3 hours

const SECRECT_KEY = "secrectkey123"

func GenerateToken(email string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: TOKEN_EXP,
		},
		UserEmail: email,
	})
	token, err := t.SignedString([]byte(SECRECT_KEY))
	return token, err
}

func TokenIsValid(c *gin.Context) bool {
	tokenStr := getTokenFromRequest(c)
	if tokenStr == "" {
		return false;
	}
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRECT_KEY), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
