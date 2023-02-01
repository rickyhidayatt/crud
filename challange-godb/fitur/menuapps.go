package fitur

import (
	"challangge-godb/controller"
	"fmt"
)

var CheckID int

func MenuApps() {
	fmt.Println("Masuk Sebagai :\n\n1.Kasir\n2.Customer\n ")
	var ask int

	for {
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&ask)

		if ask == 1 {
			fmt.Println("Kamu masuk sebagai kasir")
			KasirApps()
			break

		} else if ask == 2 {
			PrintOrderList()
			fmt.Println("Silahkan Login dengan menggunakan No ID kamu : ")
			fmt.Scan(&CheckID)

			/*validasi apakah Id Customer ada di tabel transaksi
			(sudah melakukan transaksi dan tinggal pembayaran kalo ga ada berarti customer harus minta update ke kasir)
			*/
			chekid, err := controller.CheckCustomerDB(CheckID)
			if !chekid {
				fmt.Println(err, "ID yang kamu masukan belum melakukan transaksi silahkan hubungi kasir untuk mengupdate data transaksi kamu")
			}

			controller.BillOrderDetail(CheckID)

			break

		} else {
			fmt.Println("Input yang anda masukkan salah ")
			break
		}
	}
}
