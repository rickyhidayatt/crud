package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"database/sql"
)

// func buat nampilin tabel transaksi dari DB
func GetTransactionList() []entity.Transaksi {
	db := config.ConnectDB()
	defer db.Close()

	sqlCode := "SELECT * FROM transaksi ORDER BY id;"
	rows, err := db.Query(sqlCode)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	listTRX := scanTransaction(rows)
	return listTRX
}

func scanTransaction(rows *sql.Rows) []entity.Transaksi {
	listTrx := []entity.Transaksi{}

	var err error
	for rows.Next() {
		trx := entity.Transaksi{}
		err := rows.Scan(&trx.Id, &trx.IdOrder, &trx.IdMenu, &trx.DateIn, &trx.DateOut, &trx.JumlahOrder)

		if err != nil {
			panic(err)
		}
		listTrx = append(listTrx, trx)
	}
	if err != nil {
		panic(err)
	}
	return listTrx
}

// func buat nampilin  order id daru tabel transaksi dari DB
func GetIDorderFromTx(idorder int) []entity.Transaksi {
	db := config.ConnectDB()
	defer db.Close()

	sqlCode := "select *from transaksi WHERE id_order =$1"
	rows, err := db.Query(sqlCode, idorder)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	listTRX := scanTransaction(rows)
	return listTRX
}
