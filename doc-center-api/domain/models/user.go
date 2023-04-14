package models

import "time"

type User struct {
	Id            int64  `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	NaturalPerson bool   `json:"naturalPerson"`
	CPF           string `json:"cpf"`
	// CPF           string    `json:"cpf" gorm:"check:check_cpf, naturalPerson = 1"`
	BornDate time.Time `json:"bornDate"`
	CNPJ     string    `json:"cnpj"`
	// CNPJ          string    `json:"cnpj" gorm:"check:check_cnpj, naturalPerson = 1"`
	CompanyName string `json:"companyName"`
	TradingName string `json:"tradingName"`
	Email       string `json:"email"`
	Senha       string `json:"senha"`
}

func (b *User) User() string {
	return "user"
}
