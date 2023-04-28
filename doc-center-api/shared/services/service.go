package services

import (
	"golang.org/x/crypto/bcrypt"
)

<<<<<<< HEAD
func HashPassword(password *string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
=======
func HashPassword(user *models.User) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Senha), bcrypt.DefaultCost)
>>>>>>> main

	if err != nil {
		return err.Error()
	}
<<<<<<< HEAD
	*password = string(hashedPassword)
=======
	user.Senha = string(hashedPassword)
>>>>>>> main

	return user.Senha
}
