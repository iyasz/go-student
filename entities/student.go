package entities

import "time"

type Student struct {
	Id        uint
	Nama      string
	Nis       string
	Telp      string
	CreatedAt time.Time
}
