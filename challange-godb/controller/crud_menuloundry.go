package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"fmt"
)

// Add menu
func AddMenu(m entity.Menu) {
	db := config.ConnectDB()
	defer db.Close()
	var err error

	sql := "INSERT INTO list_menu (id_menu, jenis_layanan, satuan, harga) VALUES ($1,$2,$3,$4)"

	db.Exec(sql, m.IdMenu, m.JenisLayanan, m.Satuan, m.Harga)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Berhasil menambahkan layanan baru !")
	}
}

// Update menu
func UpdateMenu(m entity.Menu) {
	db := config.ConnectDB()
	defer db.Close()
	var err error

	sql := "UPDATE list_menu SET jenis_layanan =$2, satuan =$3, harga= $4 WHERE id_menu = $1"

	db.Exec(sql, m.IdMenu, m.JenisLayanan, m.Satuan, m.Harga)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Berhasil mengupdate menu!")
	}
}

// delete menu
func Deletemenu(id int) {
	db := config.ConnectDB()
	defer db.Close()
	var err error

	sql := "DELETE FROM list_menu WHERE id_menu = $1"

	db.Exec(sql, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Berhasil menghapus menu!")
	}
}
