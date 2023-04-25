package auth

import (
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
	"doc-center-api/shared/services"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AuthService struct {
}

// Essa função pega o usuário e senha
func (s *AuthService) CreateToken(Email string, Senha string) (string, error) {

	var user models.User
	dbConfig := database.BuildDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Trate o erro aqui
		return "", err
	}

	result := db.Table("Usuario").Where("Email = ?", Email).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	user.Password = Senha
	var senhaCodificada = services.HashPassword(&user)

	if user.Id != 0 {
		if user.Password == senhaCodificada {

			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["Email"] = Email
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
			t, err := token.SignedString([]byte("secret-key"))

			//defer db.DB().Close()
			//Se der erro, retorna o erro
			if err != nil {
				return "", err
			}

			//Se não der erro, retorna o token
			return t, nil
		} else {
			//Se o usuário ou senha estiverem incorretos, retorna:
			return "", fmt.Errorf("authentication failed for user %s", Email)
		}
	}

	return "", nil
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

	/*router.GET("/verify", func(c *gin.Context) {

		//Busca o token nos cookies da requisição
		cookie, erro := c.Cookie("token")

		//Se der erro, retorna:
		if erro != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token not found",
			})
			return
		}

		//Chama a função para decodificar o token
		token, erro := authService.VerifyToken(cookie)

		//Se der erro, retorna
		if erro != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		}

		//Se der certo, retorna o nome do usuário
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": token,
		})

	})*/
}
