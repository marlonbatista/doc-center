package controllers

import (
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

func GetUserById(c *gin.Context) {
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
