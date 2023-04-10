package auth

import (
	"doc-center-api/domain/models"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	Users []models.User
}

// Essa função pega o usuário e senha
func (s *AuthService) CreateToken(UserName string, Password string) (string, error) {

	//Verificar se as credenciais estão corretas
	if UserName == "1234" && Password == "123" {

		//Codifica as credenciais e cria o token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["UserName"] = UserName
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		t, err := token.SignedString([]byte("secret-key"))

		//Se der erro, retorna o erro
		if err != nil {
			return "", err
		}

		//Se não der erro, retorna o token
		return t, nil
	} else {
		//Se o usuário ou senha estiverem incorretos, retorna:
		return "", fmt.Errorf("authentication failed for user %s", UserName)
	}
}

// Função que verifica o token
func (s *AuthService) VerifyToken(tokenString string) (string, error) {

	//Decodifica o Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret-key"), nil
	})

	//Se der erro, retorna o erro
	if err != nil {
		return "", err
	}

	//Se der certo, retorna o nome de usuário
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userName := claims["UserName"].(string)
		return userName, nil
	}

	//Se o token não for válido, retorna o erro:
	return "", errors.New("invalid token")
}
