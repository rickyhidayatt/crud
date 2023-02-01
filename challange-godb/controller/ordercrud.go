package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"fmt"
)

// ADD
func AddOrder(o entity.Order) {
	db := config.ConnectDB()
	defer db.Close()
	var err error

	sql := "INSERT INTO mst_order (id_order, nama_customer, nama_kasir, no_hp_customer, total_tagihan) VALUES ($1,$2,$3,$4,$5)"

	db.Exec(sql, o.IdOrder, o.NamaCustomer, o.NamaKasir, o.NoHP, o.TotalTagihan)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Berhasil menambahkan data orderan !")
	}
}

// Update
func UpdateOrder(o entity.Order) {
	db := config.ConnectDB()
	defer db.Close()
	var err error

	sql := "UPDATE mst_order SET nama_customer =$2, nama_kasir =$3,  no_hp_customer= $4, total_tagihan =$5  WHERE id_order = $1"

	db.Exec(sql, o.IdOrder, o.NamaCustomer, o.NamaKasir, o.NoHP, o.TotalTagihan)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Berhasil mengupdate data orderan !")
	}
}

// delete
func DeleteOrder(id int) {
	db := config.ConnectDB()
	defer db.Close()
	var err error

	sql := "DELETE FROM mst_order WHERE id_order = $1"

	db.Exec(sql, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Berhasil menghapus data orderan !")
	}
}

// GET ORDER BY ID
func ReadOrderById(id int, Bnykorder int) entity.Order {
	db := config.ConnectDB()
	defer db.Close()
	var err error

	sql := "SELECT *FROM mst_order WHERE id_order = $1 "
	read := entity.Order{}

	err = db.QueryRow(sql, id).Scan(&read.IdOrder, &read.NamaCustomer, &read.NamaKasir, &read.NoHP, &read.TotalTagihan)

	if err != nil {
		panic(err)
	}

	jumlahPesanan := entity.Transaksi{
		JumlahOrder: Bnykorder,
	}

	fmt.Println("+----ID------+------Nama------+------Kasir-----+-----No HP-----+--Jumlah Pesanan--+-----Tagihan----+")
	fmt.Printf("| %-10d | %-14s | %-14s | %-13s | %-14d | %-14d\n", read.IdOrder, read.NamaCustomer, read.NamaKasir, read.NoHP, jumlahPesanan.JumlahOrder, read.TotalTagihan)
	fmt.Println("+------------+----------------+----------------+---------------+----------------+------------------+")

	return read

}
