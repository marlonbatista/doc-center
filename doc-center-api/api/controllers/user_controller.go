package controllers

import (
	"doc-center-api/api/auth"
	"doc-center-api/domain/handlers"
	"doc-center-api/domain/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var user []models.User
	err := handlers.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// CreateUser ... Create User
func CreateUser(c *gin.Context) {
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

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := handlers.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := handlers.GetUserByID(&user, id)
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

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := handlers.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

func Login(c *gin.Context) {
	authService := auth.AuthService{}

	var loginData struct {
		Username string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	if erro := c.ShouldBindJSON(&loginData); erro != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": erro.Error(),
		})
		return
	}

	tokenString, erro := authService.CreateToken(loginData.Username, loginData.Password)

	if erro != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": erro.Error(),
		})
		return
	}

	cookie := http.Cookie{
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(c.Writer, &cookie)

	http.SetCookie(c.Writer, &cookie)

}
