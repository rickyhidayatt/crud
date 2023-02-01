package controller

import (
	"challangge-godb/config"
	"database/sql"
	"fmt"
)

func DeleteTransaksi(idorder int) {
	db := config.ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	deleteorderTX(idorder, tx)
	deleteorderOD(idorder, tx)

	err = tx.Commit()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sukses menghapus semua data order dan transaksi")
	}

}

func deleteValidasi(err error, msg string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println("transaksi delete di Rollback", err)
	} else {
		fmt.Println("Sukses", msg)
	}
}

// delet di tabel order
func deleteorderOD(id int, tx *sql.Tx) {
	sql := "DELETE FROM mst_order WHERE id_order = $1"
	_, err := tx.Exec(sql, id)
	deleteValidasi(err, "Hapus data dari tabel mst order", tx)
}

// delete di tabel transaksi
func deleteorderTX(id int, tx *sql.Tx) {
	sql := "DELETE FROM transaksi WHERE id_order = $1"
	_, err := tx.Exec(sql, id)
	deleteValidasi(err, "Hapus data dari tabel transaksi", tx)
}
