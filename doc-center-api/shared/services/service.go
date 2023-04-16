package services

import (
	"doc-center-api/domain/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return nil
}
