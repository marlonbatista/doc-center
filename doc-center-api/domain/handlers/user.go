package handlers

import (
	"fmt"

	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
	"doc-center-api/shared/services"
)

func CheckEmailAndPasswordUser(email string, password string) (err error) {
	var user models.User
	if err = database.DB.Where("email = ? and password = ?", email, password).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(user *[]models.User) (err error) {
	if err = database.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *models.User) (err error) {
	services.HashPassword(user)
	if err = database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *models.User, id string) (err error) {
	if err = database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, id string) (err error) {
	fmt.Print(user)
	database.DB.Save(user)
	return nil
}
