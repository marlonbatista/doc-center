package main

import (
	"doc-center-api/api/routes"
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {

	database.DB, err = gorm.Open(mysql.Open(database.DdURL(database.BuildDBConfig())))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer database.CloseConn()
	database.DB.AutoMigrate(&models.PermissionType{}, &models.Permission{}, &models.User{},
		&models.Dependent{}, &models.File{})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := routes.ConfigRotas()
	router.Run(":" + port)

}
