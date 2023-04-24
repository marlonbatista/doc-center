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

func GetPermissionById(c *gin.Context) {
	id := c.Params.ByName("id")
	var permission models.Permission
	err := handlers.GetPermissionById(&permission, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, permission)
	}
}
