package models

import "time"

type Permissao struct {
	Id        int64     `json:"id"`
	FileId    int64     `json:"file"`
	Ownerid   int64     `json:"owner-user"`
	GuestId   int64     `json:"guest-user"`
	Period    time.Time `json:"period"`
	Permanent bool      `json:"permanent"`
}
