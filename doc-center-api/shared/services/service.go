package services

import (
	"doc-center-api/domain/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(user *models.User) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Senha), bcrypt.DefaultCost)

	if err != nil {
		return err.Error()
	}
	user.Senha = string(hashedPassword)

	return user.Senha
}
