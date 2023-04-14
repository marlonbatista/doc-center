package main

import (
	"doc-center-api/api/routers"
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
	"fmt"

	// "github.com/go-sql-driver/mysql"
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
	database.DB.AutoMigrate(&models.User{})

	router := routers.ConfigRotas()

	//Inicializa o framewrok de rotas
	router.Run()

}
