package models

import "time"

type User struct {
	Id            int64     `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	CPF           string    `json:"cpf" gorm:"check_cpf, pessoafisica = 1"`
	BornDate      time.Time `json:"bornDate"`
	CNPJ          string    `json:"cnpj" gorm:"check: check_copy, pessoafisica = 0"`
	NaturalPerson bool      `json:"naturalPerson"`
	CompanyName   string    `json:"companyName"`
	TradingName   string    `json:"tradingName"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
}
