package fitur

import (
	"challangge-godb/controller"
	"fmt"
	"strings"
)

var NamaKasir string
var idCustomer int
var customerName string
var noHPCustomer string
var layananloundry int
var jumlahPesanan int

func KasirApps() {

	fmt.Println("Silahkan Masukan Password Kasir Karyawan ENIGMA :\n ")

	var cekpass string
	fmt.Scan(&cekpass)
	Login, _ := Login(cekpass)
	if Login {
		fmt.Println("Login berhasil")
	}
	fmt.Println(strings.Repeat("\n-", 10))

	fmt.Println("Masukan Nama Petugas :")
	fmt.Scan(&NamaKasir)

	fmt.Println("\nSilahkan Pilih menunya ", NamaKasir, ` :
	
1. Input Customer Baru
2. Tambah Order Customer
3. Delete Data Customer
4. Tambah Service Layanan Baru
5. Update Menu Layanan
6. Delete Menu Layanan
	`)

	var choice int
	fmt.Scan(&choice)

	for {

		if choice == 1 {
			KasirAddCustomer()
			break
		} else if choice == 2 {
			CustomerAddService()
			break

		} else if choice == 3 {

			fmt.Println("\nMemilih menu ini akan menghapus semua data transaksi, termasuk data customer.\nApakah kamu yakin ingin menghapus semua data ?\n\n-Ketik Y jika ingin tetap menghapus\n-Ketik N jika tidak ingin menghapus data")

			var ask string
			fmt.Scan(&ask)

			if ask == "y" || ask == "Y" {
				fmt.Println("Pilih ID ORDER yang akan di hapus :")
				var id int
				fmt.Scanln(&id)

				controller.DeleteTransaksi(id)
				fmt.Println("semua data berhasil di hapus")
				controller.ListDetailTx()

			} else if ask == "n" || ask == "N" {
				fmt.Println("")
				fmt.Println("Hallo selamat datang di Aplikasi Enigma")
				KasirApps()
			} else {
				fmt.Println("Input yang kamu masukan salah, kamu harus login kembali")
			}
			controller.ListDetailTx()

			fmt.Println("silahkan masukan ID ORDER yang ingin di hapus (ID ORDER Kolom ke 2) :")

			var idorderhapus int
			fmt.Scan(&idorderhapus)

			controller.DeleteTransaksi(idorderhapus)
			controller.ListDetailTx()

			break
		} else if choice == 4 {
			AddMenuService()
			break
		} else if choice == 5 {
			UpdateMenuService()
			break
		} else if choice == 6 {
			DeleteMenuService()
			break
		} else {
			fmt.Println("")
			fmt.Println("Pilihanmu tidak sesuai", "\nSilahkan Pilih menunya ", NamaKasir, `:
	
1. Input Customer Baru
2. Update Data Customer
3. Delete Data Customer
4. Tambah Service Layanan Baru
5. Update Menu Layanan
6. Delete Menu Layanan
				`)

			var choice int
			fmt.Scan(&choice)
			break
		}
	}

}
