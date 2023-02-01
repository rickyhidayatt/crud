package entity

import (
	"time"
)

type Transaksi struct {
	Id          int
	IdOrder     int
	IdMenu      int
	DateIn      time.Time
	DateOut     time.Time
	JumlahOrder int
}

type Order struct {
	IdOrder      int
	NamaCustomer string
	NamaKasir    string
	NoHP         string
	TotalTagihan int
}

type Menu struct {
	IdMenu       int
	JenisLayanan string
	Satuan       string
	Harga        int
}
