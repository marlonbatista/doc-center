package models

import "time"

type Dependent struct {
	Id            int64     `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name"`
	BornDate      time.Time `json:"bornDate"`
	ResponsableId int64     `json:"responsableId"`

	User User `gorm:"foreignKey:ResponsableId"`
}
