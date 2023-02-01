package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"fmt"
	"strings"
)

func totalorder(idorder int, namacust string, namakasir string, nohp string) {
	db := config.ConnectDB()
	defer db.Close()

	sql := `
SELECT transaksi.id_order, mst_order.total_tagihan
FROM mst_order
JOIN transaksi ON mst_order.id_order = transaksi.id_order
JOIN list_menu ON transaksi.id_menu = list_menu.id_menu
WHERE transaksi.id_order = $1
LIMIT 1
`
	rows, err := db.Query(sql, idorder)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	order := entity.Order{}

	transaksi := entity.Transaksi{}

	for rows.Next() {
		if err := rows.Scan(&transaksi.Id, &order.TotalTagihan); err != nil {
			fmt.Println("Total Order, Query salah")
			panic(err)
		}

		fmt.Println(strings.Repeat("=", 50))
		fmt.Println("")
		fmt.Println("Nama Customer :", strings.ToUpper(namacust))
		fmt.Println("No HP Customer :", nohp)
		fmt.Println("Kasir :", strings.ToUpper(namakasir))
		fmt.Println(strings.ToUpper(namacust), "Berikut total tagihan kamu Rp.", order.TotalTagihan)
		fmt.Println("")
		fmt.Println(strings.Repeat("=", 50))
	}
}

// Dapatkan semua data orderan by ID order
func BillOrderDetail(idorder int) {
	db := config.ConnectDB()
	defer db.Close()

	sql := `
	SELECT transaksi.id,mst_order.id_order, transaksi.tanggal_masuk, transaksi.tanggal_keluar, mst_order.nama_kasir, mst_order.nama_customer,list_menu.jenis_layanan, transaksi.jumlah_order, list_menu.satuan, list_menu.harga, mst_order.no_hp_customer
	FROM mst_order
    JOIN transaksi ON mst_order.id_order = transaksi.id_order
	JOIN list_menu ON transaksi.id_menu = list_menu.id_menu
	WHERE transaksi.id_order = $1;
	`

	rows, err := db.Query(sql, idorder)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	order := entity.Order{}
	menu := entity.Menu{}
	trx := entity.Transaksi{}

	fmt.Println(strings.Repeat("=", 32))
	fmt.Println("|  BILL DETAIL LOUNDRY ENIGMA  |")
	fmt.Println(strings.Repeat("=", 32))

	fmt.Println("")

	fmt.Printf("%-5s | %-10s | %-15s | %-15s | %-15s | %-15s | %-15s | %-15s \n", "No", "ID Order", "Tanggal Masuk", "Tanggal Keluar", "Jenis Layanan", "Jumlah Order", "Satuan", "Harga/Layanan")
	fmt.Println(strings.Repeat("=", 125))

	for rows.Next() {

		err := rows.Scan(&trx.Id, &trx.IdOrder, &trx.DateIn, &trx.DateOut, &order.NamaKasir, &order.NamaCustomer, &menu.JenisLayanan, &trx.JumlahOrder, &menu.Satuan, &menu.Harga, &order.NoHP)
		if err != nil {
			panic(err)
		}

		fmt.Printf("\n%-5d | %-10d | %-15s | %-15s | %-15s | %-15d | %-15s | %-15d \n", trx.Id, trx.IdOrder, trx.DateIn.Format("2006-01-02"), trx.DateOut.Format("2006-01-02"), menu.JenisLayanan, trx.JumlahOrder, menu.Satuan, menu.Harga)

	}
	totalorder(idorder, order.NamaCustomer, order.NamaKasir, order.NoHP)
}
