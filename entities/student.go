package entities

import "time"

type Student struct {
	Id        uint
	Nama      string
	Nis       string
	Umur      string
	Gender    string
	Telp      string
	Alamat    string
	CreatedAt time.Time
}
