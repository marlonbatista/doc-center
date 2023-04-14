package main

import (
	"doc-center-api/api/auth"
	"doc-center-api/api/routers"
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {

	database.DB, err = gorm.Open(mysql.Open(database.DdURL(database.BuildDBConfig())))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer database.CloseConn()
	database.DB.AutoMigrate(&models.PermissionType{}, &models.Permission{}, &models.User{},
		&models.Dependent{}, &models.File{})

	router := routers.ConfigRotas()

	//Instancia a model de usuário
	authService := auth.AuthService{
		Users: []models.User{},
	}

	//Instancia o router; função que manipula as rotas HTTP
	// router := gin.Default()

	// router.POST("/login", func(responseHttp *gin.Context) { //O responseHttp é a variável que vai armazenar os dados enviados pelo Http

	// 	//Cria uma struct consumindo os dados do
	// 	var loginData struct {
	// 		Username string `json:"username" binding:"required"`
	// 		Password string `json:"password" binding:"required"`
	// 	}

	// 	//Chama a função ShouldBindJSON que verifica se o loginData é válido
	// 	if erro := responseHttp.ShouldBindJSON(&loginData); erro != nil {

	// 		responseHttp.JSON(http.StatusBadRequest, gin.H{
	// 			"error": erro.Error(),
	// 		})
	// 		return
	// 	}

	// 	//Chama a função CreateToken e gera o token do usuário
	// 	tokenString, erro := authService.CreateToken(loginData.Username, loginData.Password)

	// 	//Se der erro no token
	// 	if erro != nil {
	// 		responseHttp.JSON(http.StatusUnauthorized, gin.H{
	// 			"error": erro.Error(),
	// 		})
	// 		return
	// 	}

	// 	cookie := http.Cookie{
	// 		Name:     "token",
	// 		Value:    tokenString,
	// 		HttpOnly: true,
	// 		Secure:   true,
	// 		Path:     "/",
	// 	}

	// 	//Insere o cookie com o token gerado
	// 	http.SetCookie(responseHttp.Writer, &cookie)

	// 	//Retorna uma mensagem
	// 	responseHttp.JSON(http.StatusOK, gin.H{"message": "login successful"})

	// })

	router.GET("/verify", func(c *gin.Context) {

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

	})

	//Inicializa o framewrok de rotas
	router.Run()

}
