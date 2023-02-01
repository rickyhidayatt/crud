package fitur

import (
	"challangge-godb/controller"
	"challangge-godb/entity"
	"fmt"
	"time"
)

func CustomerAddService() {

	PrintOrderList()
	fmt.Println("")
	fmt.Println("Pilih ID Customer yang mau di tambah orderan :")
	fmt.Scan(&idCustomer)

	chekid, err := controller.CheckCustomerDB(idCustomer)
	if !chekid {
		fmt.Println(err, "ID Yang kamu  masukan salah")
		CustomerAddService()
	}

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

	fmt.Println("Update berapa banyak pesanan : ")
	fmt.Scan(&jumlahPesanan)

	PrintTransactionList()
	fmt.Println("masukan NO ID Baru :")
	fmt.Scan(&IDTRX)

	tambahorder := entity.Transaksi{
		Id:          IDTRX,
		IdOrder:     idCustomer,
		IdMenu:      layananloundry,
		DateIn:      time.Now(),
		DateOut:     time.Now().Add(time.Hour * 24),
		JumlahOrder: jumlahPesanan,
	}

	controller.Transaksi(tambahorder)

	controller.BillOrderDetail(tambahorder.IdOrder)
}
