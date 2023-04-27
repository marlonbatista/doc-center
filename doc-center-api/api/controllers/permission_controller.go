package controllers

import (
	"doc-center-api/domain/handlers"
	"doc-center-api/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPermissions(c *gin.Context) {
	var permission []models.Permission
	err := handlers.GetAllPermissions(&permission)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, permission)
	}
}

func GetOwnerPermission(c *gin.Context) {
	id := c.Params.ByName("id")
	var permission []models.Permission
	err := handlers.GetOwnerPermission(id, &permission)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, permission)
	}
}

func GetGuestPermission(c *gin.Context) {
	id := c.Params.ByName("id")
	var permission []models.Permission
	err := handlers.GetGuestPermission(id, &permission)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, permission)
	}
}

func CreatePermission(c *gin.Context) {
	var permission models.Permission

	err := c.BindJSON(&permission)
	if err != nil {
		return
	}

	err = handlers.CreatePermission(&permission)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, permission)
	}
}

func UpdatePermission(c *gin.Context) {
	var permission models.Permission
	err := c.BindJSON(&permission)
	if err != nil {
		return
	}

	err = handlers.UpdatePermission(&permission)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, permission)
	}
}

func DeletePermission(c *gin.Context) {
	id := c.Params.ByName("id")
	var permission models.Permission
	err := handlers.DeletePermission(id, &permission)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, permission)
	}
}
