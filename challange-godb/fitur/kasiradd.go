package fitur

import (
	"challangge-godb/controller"
	"challangge-godb/entity"
	"fmt"
	"strings"
	"time"
)

var IDTRX int

func KasirAddCustomer() {

	PrintOrderList()
	fmt.Println("")
	fmt.Println("Masukan ID Customer Baru :")
	fmt.Scan(&idCustomer)

	fmt.Println("")
	fmt.Println("Masukan Nama Customer :")
	fmt.Scan(&customerName)

	fmt.Println("")
	fmt.Println("Masukan No HP Customer : ")
	fmt.Scan(&noHPCustomer)

	fmt.Println("")
	fmt.Println("Pilih Layanan Loundry")

	//panggil menu londry
	PrintListMenu()
	//end
	fmt.Println("")
	fmt.Println("Mau Pilih Layanan Nomor Berapa ? :")
	fmt.Scan(&layananloundry)

	// validasi jumlah id di database menu
	count, err := controller.GetIDCount()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Jumlah layanan kami hanya ada :", count)
	}

	for layananloundry > count {
		fmt.Println("Silahkan ulangi pilihan kamu :")
		fmt.Scan(&layananloundry)
	}

	fmt.Println("Mau berapa banyak ? : ")
	fmt.Scan(&jumlahPesanan)

	ADDOrder := entity.Order{
		IdOrder:      idCustomer,
		NamaCustomer: customerName,
		NoHP:         noHPCustomer,
		NamaKasir:    NamaKasir,
		TotalTagihan: 0,
	}

	jumlahorder := entity.Transaksi{
		JumlahOrder: jumlahPesanan,
	}

	controller.AddOrder(ADDOrder)
	controller.ReadOrderById(idCustomer, jumlahorder.JumlahOrder)

	fmt.Print("\nSudah sesuai atau ada perubahan ? \n-Ketik Y untuk update\n-Ketik D untuk Delete \n-Ketik N untuk lanjutkan program \n")

	var YorN string
	fmt.Scan(&YorN)

	if YorN == "y" || YorN == "Y" {
		KasirUpdate()
	} else if YorN == "n" || YorN == "N" {

		// list detail transaksi dari 2 tabel (order dan transaksi)
		controller.ListDetailTx()

		fmt.Println("Masukan ID Transaksi baru :")
		fmt.Scan(&IDTRX)

		transaksi := entity.Transaksi{
			Id:          IDTRX,
			IdOrder:     idCustomer,
			IdMenu:      layananloundry,
			DateIn:      time.Now(),
			DateOut:     time.Now().Add(time.Hour * 24),
			JumlahOrder: jumlahPesanan,
		}

		controller.Transaksi(transaksi)
		// controller.BillDetail(IDTRX)
		controller.BillOrderDetail(transaksi.IdOrder)

	} else if YorN == "d" || YorN == "D" {
		KasiDelete()
		fmt.Println("Data Customer Berhasil Terhapus")
		fmt.Print(strings.Repeat("\n-", 3))

		fmt.Println("Hallo Selamat datang kembali di Aplikasi ENIGMA LOUNDRY")
		fmt.Println("")
		KasirApps()
	} else {
		fmt.Println("Data tidak valid \nJika anda ingin update silahkan jalankan ulang program")
	}

}
