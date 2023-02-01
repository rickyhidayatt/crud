package fitur

import (
	"challangge-godb/controller"
	"challangge-godb/entity"
	"fmt"
	"strings"
	"time"
)

func KasirUpdate() {

	PrintOrderList()
	fmt.Println("")
	fmt.Println("Pilih ID Customer yang mau di Update :")
	fmt.Scan(&idCustomer)

	fmt.Println("")
	fmt.Println("Update Nama Customer :")
	fmt.Scan(&customerName)

	fmt.Println("")
	fmt.Println("Update No HP Customer : ")
	fmt.Scan(&noHPCustomer)

	fmt.Println("")
	fmt.Println("Update Layanan Loundry")

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

	fmt.Println("Update berapa banyak pesanan : ")
	fmt.Scan(&jumlahPesanan)

	UpdateOrder := entity.Order{
		IdOrder:      idCustomer,
		NamaCustomer: customerName,
		NoHP:         noHPCustomer,
		NamaKasir:    NamaKasir,
		TotalTagihan: 0,
	}

	Jmlahorder := entity.Transaksi{
		JumlahOrder: jumlahPesanan,
	}

	controller.UpdateOrder(UpdateOrder)
	controller.ReadOrderById(idCustomer, Jmlahorder.JumlahOrder)

	fmt.Print("\nSudah sesuai atau ada perubahan lagi ? \n-Ketik Y untuk update lagi\n-Ketik D untuk Delete \n-Ketik N untuk lanjutkan program\n")

	//kasir bisa update dan delete data dari sini kalo semisal data yg di input salah

	var YorN string
	fmt.Scan(&YorN)

	if YorN == "y" || YorN == "Y" {
		KasirUpdate()
	} else if YorN == "n" || YorN == "N" {
		fmt.Println("Input Sukses")
	} else if YorN == "d" || YorN == "D" {
		KasiDelete()
		fmt.Print(strings.Repeat("\n-", 5))
		fmt.Println("Hallo Selamat datang kembali di Aplikasi ENIGMA LOUNDRY")
		fmt.Println("")
		KasirApps()
	} else {
		fmt.Println("Data tidak valid \nJika anda ingin update silahkan jalankan ulang program")
	}

	// list detail transaksi dari 2 tabel (order dan transaksi)
	controller.ListDetailTx()

	fmt.Println("Masukan ID Transaksi baru untuk di Update:")
	fmt.Scan(&IDTRX)

	transaksi := entity.Transaksi{
		Id:          IDTRX,
		IdOrder:     idCustomer,
		IdMenu:      layananloundry,
		DateIn:      time.Now(),
		DateOut:     time.Now().Add(time.Hour * 4),
		JumlahOrder: jumlahPesanan,
	}

	controller.Transaksi(transaksi)

	controller.BillOrderDetail(transaksi.IdOrder)
}
