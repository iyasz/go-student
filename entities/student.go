package entities

import "time"

type Student struct {
	Id        uint
	Nama      string
	Nis       uint32
	Telp      string
	CreatedAt time.Time
}
