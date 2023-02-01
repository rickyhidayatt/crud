package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"fmt"
	"strings"
)

//menampilkan semua data tabel transakasi dan tabel order

func ListDetailTx() {
	db := config.ConnectDB()
	defer db.Close()

	sql := `
	SELECT transaksi.id AS id_transaksi, transaksi.id_order, mst_order.nama_customer, mst_order.no_hp_customer, transaksi.jumlah_order, mst_order.total_tagihan
	FROM transaksi
	JOIN mst_order ON transaksi.id_order = mst_order.id_order
	ORDER BY id;
	`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println(strings.Repeat("+", 20))
	fmt.Printf("\n |%-15s| %-15s| %-15s| %-15s| %-15s| %-15s| \n", "ID TRANSAKSI", "ID ORDER", "NAMA CUSTOMER", "NO HP", "JUMLAH ORDER", "TAGIHAN")
	fmt.Print(strings.Repeat("-", 103))

	for rows.Next() {
		transaksi := entity.Transaksi{}
		order := entity.Order{}

		err := rows.Scan(&transaksi.Id, &transaksi.IdOrder, &order.NamaCustomer, &order.NoHP, &transaksi.JumlahOrder, &order.TotalTagihan)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n |%-15d| %-15d| %-15s| %-15s| %-15d| %-15d|\n", transaksi.Id, transaksi.IdOrder, order.NamaCustomer, order.NoHP, transaksi.JumlahOrder, order.TotalTagihan)
	}

}
