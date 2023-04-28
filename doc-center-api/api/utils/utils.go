package utils

import (
	// "errors"
	"fmt"
	// "log"
	"os"
	// "strconv"
	"strings"
	"time"

	// "doc-center-api/domain/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.StandardClaims
	UserEmail string
}

var TOKEN_EXP = int64(time.Hour * 3) // 3 hours

var privateKey = []byte(os.Getenv("API_SECRET"))

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
	// tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	// if err != nil {
	// 	return "", err
	// }

	// claims := jwt.MapClaims{}
	// claims["authorized"] = true
	// claims["id"] = user.Id
	// claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// return token.SignedString(privateKey)

}

func TokenIsValid(c *gin.Context) bool {

	tokenStr := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRECT_KEY), nil
	})
	if err != nil {
		return false
	}
	return token.Valid
	// token, err := GetToken(c)

	// if err != nil {
	// 	return err
	// }

	// _, ok := token.Claims.(jwt.MapClaims)
	// if ok && token.Valid {
	// 	return nil
	// }

	// return errors.New("Invalid token provided")
}

// func CurrentUser(c *gin.Context) (models.User, error) {
// 	err := ValidateToken(c)
// 	if err != nil {
// 		return models.User{}, err
// 	}
// 	token, _ := GetToken(c)
// 	claims, _ := token.Claims.(jwt.MapClaims)
// 	userId := int(claims["id"].(float64))
// 	var user models.User
// 	erro := models.GetUserById(&user, userId)
// 	if erro != nil {
// 		return models.User{}, err
// 	}
// 	return user, nil
// }

func GetToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")

	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
