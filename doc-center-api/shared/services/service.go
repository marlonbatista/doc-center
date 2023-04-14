package services

import (
	"crypto/sha256"
	"fmt"

	"doc-center-api/domain/models"

	"golang.org/x/crypto/bcrypt"
)

func SHA256Encoder(s string) string {
	str := sha256.Sum256([]byte(s))
	// bcrypt
	return fmt.Sprintf("%x", str)
}

func HashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return nil
}
