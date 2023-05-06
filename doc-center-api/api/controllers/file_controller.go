package controllers

import (
	"doc-center-api/domain/handlers"
	"doc-center-api/domain/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllFiles(c *gin.Context) {
	var arquivo []models.File
	err := handlers.GetAllFiles(&arquivo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, arquivo)
	}
}

func GetFileById(c *gin.Context) {
	id := c.Params.ByName("id")
	var file models.File
	err := handlers.GetFileByID(&file, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, file)
	}
}

func CreateFile(c *gin.Context) {
	var file models.File
	err := c.BindJSON(&file)
	if err != nil {
		return
	}

	err = handlers.CreateFile(&file)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, file)
	}
}

func UpdateFile(c *gin.Context) {
	var file models.File
	err := c.BindJSON(&file)
	if err != nil {
		return
	}

	err = handlers.UpdateFile(&file)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, file)
	}
}

func DeleteFile(c *gin.Context) {
	var file models.File
	id := c.Params.ByName("id")

	err := handlers.DeleteFile(&file, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, file)
	}
}
