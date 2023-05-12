package controllers

import (
	"doc-center-api/domain/handlers"
	"doc-center-api/domain/models"
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
	us, err := handlers.GetUserByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, us)
	}
}

func GetUserByName(c *gin.Context) {
	name := c.Params.ByName("name")
	var user []models.User
	err := handlers.GetUserByName(&user, name)
	if err != nil || len(user) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"User id not found: ": idParam})
		return
	}
<<<<<<< HEAD
	//err = handlers.GetUserByID(&user, id)
	//if err != nil {
	//	c.JSON(http.StatusNotFound, user)
	//}
	c.BindJSON(&user)
=======
	var user models.User
	err = nil
	if err = c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"The submitted object is not valid": err})
	}
>>>>>>> main
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
	loginReturn, err := handlers.LoginCheck(login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The e-mail or password is not correct"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":  loginReturn.Token,
		"idUser": loginReturn.IdUser,
	})

}

func GetUserPermission(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	err := handlers.GetUserPermission(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
