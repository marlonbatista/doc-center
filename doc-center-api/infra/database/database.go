package database

import (
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "125678",
		DBName:   "DocCenterDB",
	}
	return &dbConfig
}

func DdURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func CloseConn() error {
	config, err := DB.DB()
	if err != nil {
		return err
	}
	err = config.Close()
	if err != nil {
		return err
	}
	return nil
}
