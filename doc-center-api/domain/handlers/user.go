package handlers

import (
	"fmt"

	// "doc-center-api/api/auth"
	"doc-center-api/api/utils"
	// "doc-center-api/api/utils"
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
	"doc-center-api/shared/services"

	"golang.org/x/crypto/bcrypt"
)

func LoginCheck(email string, password string) (string, error) {
	var err error
	user := models.User{}
	// Model(models.User{})
	if err = database.DB.Where("email=?", email).First(&user).Error; err != nil {
		return "", err
	}

	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(email)
	return token, nil
}

func GetAllUsers(user *[]models.User) (err error) {
	if err = database.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *models.User) (err error) {
	services.HashPassword(&user.Password)
	if err = database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *models.User, id int) (err error) {
	if err = models.GetUserById(user, id); err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User, id int) (err error) {
	fmt.Print(user)
	database.DB.Save(user)
	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserPermission(user *models.User, id string) (err error) {
	if err = database.DB.InnerJoins("permissions").Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func updatePassword(user *models.User, senhaNova string, senhaAntiga string) string {
	var usuario models.User
	usuario.Senha = senhaAntiga
	var senhaCodificada = services.HashPassword(&usuario)

	if senhaCodificada != user.Senha {
		return "Senhas não conferem"
	}

	usuario.Senha = senhaNova
	senhaCodificada = services.HashPassword(&usuario)

	if database.DB.Update("Password", senhaCodificada).Where("id = ?", user.Id).First(user).Error != nil {
		return "Ocorreu um erro"
	} else {
		return "Senha alterada com sucesso"
	}
}
