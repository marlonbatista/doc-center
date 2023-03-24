package models

import "time"

type Permissao struct {
	Id              int64     `json:"id"`
	IdArquivo       int64     `json:"arquivo"`
	IdUsuarioDono   int64     `json:"usuario-dono"`
	IdUsuarioAcesso int64     `json:"usuario-acesso"`
	TempoDuracao    time.Time `json:"tempo-duracao"`
	Permante        bool      `json:"permanente"`
}
