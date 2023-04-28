package services

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	*password = string(hashedPassword)

	return nil
}
