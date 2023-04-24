package controllers

import (
	"doc-center-api/domain/handlers"
	"doc-center-api/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
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
