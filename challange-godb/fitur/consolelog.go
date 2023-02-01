package fitur

import (
	"challangge-godb/controller"
	"fmt"
	"strings"
)

func PrintTransactionList() {
	listTrx := controller.GetTransactionList()

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("ID\tID Order\tID Menu\t        Date Out\tDate In\t      Jumlah Order")
	fmt.Println(strings.Repeat("-", 80))

	for _, trx := range listTrx {
		fmt.Printf("%d\t%d\t\t%d\t\t%s\t%s\t%d\n", trx.Id, trx.IdOrder, trx.IdMenu, trx.DateIn.Format("2006-01-02"), trx.DateOut.Format("2006-01-02"), trx.JumlahOrder)
	}
}

func PrintOrderList() {
	lihatInputKasir := controller.GetOrderLoundry()
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("ID\tNama Kasir\tNama Customer\tNo HP\t   Total Tagihan")
	fmt.Println(strings.Repeat("-", 60))

	for _, data := range lihatInputKasir {
		fmt.Printf("%d\t%-10s\t%-10s\t%-10s\t%-10d\n", data.IdOrder, data.NamaKasir, data.NamaCustomer, data.NoHP, data.TotalTagihan)
	}

}

func PrintListMenu() {
	listMenu := controller.GetMenuLoundry()

	fmt.Println(strings.Repeat("-", 37))
	fmt.Println("No\tJenis Layanan\tSatuan\tHarga")
	fmt.Println(strings.Repeat("-", 37))

	for _, data := range listMenu {
		fmt.Printf("%d\t%-10s\t%-5s\t%d\n", data.IdMenu, data.JenisLayanan, data.Satuan, data.Harga)
	}
}

func PrintOrderIdfromTx(idorder int) {
	listTrx := controller.GetIDorderFromTx(idorder)

	fmt.Println(strings.Repeat("-", 80))
	fmt.Println("ID\tID Order\tID Menu\t        Date Out\tDate In\t      Jumlah Order")
	fmt.Println(strings.Repeat("-", 80))

	for _, trx := range listTrx {
		fmt.Printf("%d\t%d\t\t%d\t\t%s\t%s\t%d\n", trx.Id, trx.IdOrder, trx.IdMenu, trx.DateIn.Format("2006-01-02"), trx.DateOut.Format("2006-01-02"), trx.JumlahOrder)
	}
}
