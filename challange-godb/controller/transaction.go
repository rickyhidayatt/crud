package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"database/sql"

	"fmt"
)

func Transaksi(trf entity.Transaksi) {
	db := config.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	insertTransaction(trf, tx)
	hargatotal := totalHarga(trf.IdOrder, tx)
	updateHarga(hargatotal, trf.IdOrder, tx)

	err = tx.Commit()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Transaksi Commit")
	}

}

// Rollback func

func validasi(err error, massage string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println("transaksi Rollback", err)
	} else {
		fmt.Println("sukses", massage)
	}

}

func insertTransaction(trf entity.Transaksi, tx *sql.Tx) {
	sql := "INSERT INTO transaksi (id, id_order, id_menu, tanggal_masuk, tanggal_keluar, jumlah_order) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := tx.Exec(sql, trf.Id, trf.IdOrder, trf.IdMenu, trf.DateIn, trf.DateOut, trf.JumlahOrder)
	validasi(err, "insert transaksi", tx)

}

func totalHarga(idorder int, tx *sql.Tx) []int {
	sql := "SELECT SUM (list_menu.harga * transaksi.jumlah_order) FROM transaksi JOIN list_menu ON transaksi.id_menu = list_menu.id_menu WHERE id_order = $1"

	rows, err := tx.Query(sql, idorder)
	validasi(err, "Tambahin Harga", tx)
	defer rows.Close()

	var harga []int
	for rows.Next() {
		var h int
		err := rows.Scan(&h)
		validasi(err, "Tambahin Harga", tx)
		harga = append(harga, h)
	}
	return harga
}

func updateHarga(harga []int, id int, tx *sql.Tx) {
	var total int
	for _, h := range harga {
		total += h
	}
	sql := "UPDATE mst_order SET total_tagihan = $1 WHERE id_order = $2"
	_, err := tx.Exec(sql, total, id)

	validasi(err, "Update Harga berhasil diterapkan", tx)
}

// func totalHarga(idorder int, tx *sql.Tx) int {
// 	sql := "SELECT SUM (list_menu.harga * transaksi.jumlah_order) FROM transaksi JOIN list_menu ON transaksi.id_menu = list_menu.id_menu WHERE id_order = $1"

// 	var harga = 0
// 	err := tx.QueryRow(sql, idorder).Scan(&harga)

// 	validasi(err, "Tambahin Harga", tx)
// 	return harga
// }

// func updateHarga(harga []int, id int, tx *sql.Tx) {
// 	sql := "UPDATE mst_order SET total_tagihan = $1 WHERE id_order = $2"
// 	_, err := tx.Exec(sql, harga, id)

// 	validasi(err, "Update Harga berhasil diterapkan", tx)

// }
