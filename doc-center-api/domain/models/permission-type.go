package models

type PermissionType struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
}
