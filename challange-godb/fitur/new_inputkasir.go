package fitur

import (
	"challangge-godb/controller"
	"challangge-godb/entity"
	"fmt"
)

var idLayanan int
var namaLayanan string
var satuan string
var harga int

func AddMenuService() {

	PrintListMenu()
	fmt.Println("Masukan ID Layanan Baru :")
	fmt.Scan(&idLayanan)
	fmt.Println("Masukan Jenis Layanan Baru :")
	fmt.Scan(&namaLayanan)
	fmt.Println("Masukan Jenis Satuan :")
	fmt.Scan(&satuan)
	fmt.Println("Masukan Harga Layanan per ", satuan, ":")
	fmt.Scan(&harga)

	menubaru := entity.Menu{
		IdMenu:       idLayanan,
		JenisLayanan: namaLayanan,
		Satuan:       satuan,
		Harga:        harga,
	}

	controller.AddMenu(menubaru)
	PrintListMenu()

	fmt.Println("Menu layanan berhasil ditambahkan")

}

func DeleteMenuService() {
	PrintListMenu()
	fmt.Println("Masukan ID Layanan yang ingin di hapus :")
	fmt.Scan(&idLayanan)
	controller.Deletemenu(idLayanan)
	PrintListMenu()

	fmt.Println("Menu layanan berhasil dihapus")
}

func UpdateMenuService() {
	PrintListMenu()
	fmt.Println("Masukan ID Layanan yang ingin di Update :")
	fmt.Scan(&idLayanan)
	fmt.Println("Masukan Jenis Layanan :")
	fmt.Scan(&namaLayanan)
	fmt.Println("Masukan Jenis Satuan :")
	fmt.Scan(&satuan)
	fmt.Println("Masukan Harga Layanan per ", satuan, ":")
	fmt.Scan(&harga)

	menuupdate := entity.Menu{
		IdMenu:       idLayanan,
		JenisLayanan: namaLayanan,
		Satuan:       satuan,
		Harga:        harga,
	}

	controller.UpdateMenu(menuupdate)
	PrintListMenu()

	fmt.Println("Menu layanan berhasil di Update")
}
