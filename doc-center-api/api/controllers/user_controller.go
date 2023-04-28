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

func GetUserByID(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		panic(err)
	}
	var user models.User
	err = handlers.GetUserByID(&user, id)
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

}
