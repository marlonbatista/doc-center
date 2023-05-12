package controllers

import (
	"doc-center-api/domain/handlers"
	"doc-center-api/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllDependents(c *gin.Context) {
	var dependent []models.Dependent
	err := handlers.GetAllDependents(&dependent)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dependent)
	}
}

func GetDependentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var dependent models.Dependent
	err := handlers.GetDependentById(&dependent, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, dependent)
	}
}

func CreateDependent(c *gin.Context) {
	var dependent models.Dependent
	c.BindJSON(&dependent)
	err := handlers.CreateDependent(&dependent)
	if err != nil || dependent.Id == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, dependent)
	}
}

func DeleteDependent(c *gin.Context) {
	id := c.Params.ByName("id")
	err := handlers.DeleteDependent(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func UpdateDependent(c *gin.Context) {
	var dependent models.Dependent
	c.BindJSON(&dependent)
	err := handlers.UpdateDependent(&dependent)
	if err != nil || dependent.Id == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, dependent)
	}
}
