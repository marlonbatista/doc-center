package controllers

import (
	"doc-center-api/domain/handlers"
	"doc-center-api/domain/models"
	// "doc-center-api/shared/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	err := handlers.GetAllUsers(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func Signup(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := handlers.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

<<<<<<< HEAD
func GetUserByID(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		panic(err)
	}
=======
func GetUserById(c *gin.Context) {
	id := c.Params.ByName("id")
>>>>>>> main
	var user models.User
	err = handlers.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := handlers.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		panic(err)
	}
	err = handlers.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = handlers.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func Login(c *gin.Context) {
<<<<<<< HEAD
	var login models.Login
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := handlers.LoginCheck(login.Email, login.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The e-mail or password is not correct"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

=======
}

// imagino que este bloco de codigo pertenca a func Login
//var authService = auth.AuthService{}
//
//var loginData struct {
//	Username string `json:"Email" binding:"required"`
//	Password string `json:"Password" binding:"required"`
//}
//
//if erro := c.ShouldBindJSON(&loginData); erro != nil {
//
//	c.JSON(http.StatusBadRequest, gin.H{
//		"error": erro.Error(),
//	})
//	return
//}
//
//tokenString, erro := authService.CreateToken(loginData.Username, loginData.Password)
//
//if erro != nil {
//	c.JSON(http.StatusUnauthorized, gin.H{
//		"error": erro.Error(),
//	})
//	return
//}
//
//cookie := http.Cookie{
//	Name:     "token",
//	Value:    tokenString,
//	HttpOnly: true,
//	Secure:   true,
//	Path:     "/",
//}
//
//http.SetCookie(c.Writer, &cookie)
//
//http.SetCookie(c.Writer, &cookie)
// comentei ate aqui em cima

//func Login(c *gin.Context) {
//
//}

func GetUserPermission(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := handlers.GetUserPermission(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
>>>>>>> main
}
