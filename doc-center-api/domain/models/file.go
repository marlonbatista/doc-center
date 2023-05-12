package models

type File struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Number      string `json:number`
	DataOfIssue string `json:dataOfIssue`
	File        []byte `json:"file"`
	UserId      int64  `json:"userId"`
	User        User   `gorm:"foreignKey:UserId"`
}
