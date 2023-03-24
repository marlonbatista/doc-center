package models

import "time"

type Usuario struct {
	Id             int64     `json:"id" gorm:"primaryKey"`
	Nome           string    `json:"nome"`
	CPF            string    `json:"cpf" gorm:"check_cpf, pessoafisica = 1"`
	DataNascimento time.Time `json:"datanascimento"`
	CNPJ           string    `json:"cnpj" gorm:"check: check_cnpj, pessoafisica = 0"`
	PessoaFisica   bool      `json:"pessoafisica"`
	RazaoSocial    string    `json:"razaosocial"`
	NomeFantasia   string    `json:"nomefantasia"`
	Email          string    `json:"email"`
	Senha          string    `json:"senha"`
}
