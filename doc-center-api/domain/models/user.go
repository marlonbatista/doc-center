package models

import (
	"doc-center-api/infra/database"
)

type User struct {
	Id            int64  `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	NaturalPerson bool   `json:"naturalPerson"`
	CPF           string `json:"cpf"`
	// CPF           string    `json:"cpf" gorm:"check:check_cpf, naturalPerson = 1"`
	BornDate string `json:"bornDate"`
	CNPJ     string `json:"cnpj"`
	// CNPJ          string    `json:"cnpj" gorm:"check:check_cnpj, naturalPerson = 1"`
	CompanyName string `json:"companyName"`
	TradingName string `json:"tradingName"`
	Email       string `json:"email"`
	Senha       string `json:"password"`
}

func (b *User) User() string {
	return "user"
}

func GetUserById(user *User, id int) (err error) {
	if err = database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
