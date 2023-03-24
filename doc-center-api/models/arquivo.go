package models

type Arquivo struct {
	Id        int64  `json:"id"`
	Descricao string `json:"descricao"`
	Arquivo   []byte `json:"arquivo"`
	IdUsuario int64  `json:"usuario"`
}
