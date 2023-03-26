package models

import (
	"time"
)

type Permission struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	FileId    int64     `json:"file"`
	OwnerId   int64     `json:"ownerUser"`
	GuestId   int64     `json:"guestUser"`
	Period    time.Time `json:"period"`
	Permanent bool      `json:"permanent"`

	File      File `gorm:"foreignKey:FileId"`
	OwnerUser User `gorm:"foreignKey:OwnerId"`
	GuestUser User `gorm:"foreignKey:GuestId"`
}
