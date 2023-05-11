package handlers

import (
	"doc-center-api/api/utils"
	"doc-center-api/domain/models"
	"doc-center-api/infra/database"
	"doc-center-api/shared/services"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type loginReturn struct {
	Token  string `json:token`
	IdUser int    `json:idUser`
}

func LoginCheck(email string, password string) (loginReturn, error) {
	var err error
	user := models.User{}
	if err = database.DB.Where("email=?", email).First(&user).Error; err != nil {
		return loginReturn{}, err
	}
	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return loginReturn{}, err
	}
	token, err := utils.GenerateToken(email)
	if err != nil {
		return loginReturn{}, err
	}
	return loginReturn{
		Token:  token,
		IdUser: int(user.Id),
	}, nil
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

func GetUserByName(user *[]models.User, name string) (err error) {
	if err = database.DB.Where("name LIKE ?", "%"+name+"%").
		Limit(2).
		Offset(0).
		Find(&user).Error; err != nil {
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

func UpdatePassword(user *models.User, newPassword string, oldPassword string) bool {
	var voidUser models.User
	voidUser.Password = oldPassword
	services.HashPassword(&voidUser.Password)

	if voidUser.Password != user.Password {
		return false
	}

	voidUser.Password = newPassword
	services.HashPassword(&voidUser.Password)

	if database.DB.Update("Password", voidUser.Password).Where("id = ?", user.Id).First(user).Error != nil {
		return false
	} else {
		return true
	}
}
