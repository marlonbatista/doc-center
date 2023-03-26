package models

type File struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	File        []byte `json:"file"`
	UserId      int64  `json:"idUser"`
	User        User   `gorm:"foreignKey:UserId"`
}
