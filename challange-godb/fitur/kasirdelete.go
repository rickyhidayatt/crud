package fitur

import (
	"challangge-godb/controller"
	"fmt"
	"strconv"
)

func KasiDelete() {
	PrintOrderList()
	fmt.Println("Silahkan masukan ID yang ingin di hapus")

	var input string

	for {
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&input)

		no, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Kamu salah menginput, masukkan angka.")
		} else if no <= 0 {
			fmt.Println("Tidak ada ID yang kurang dari 1")
		} else {
			fmt.Println("Angka yang dimasukkan:", input)
			controller.DeleteOrder(no)
			break
		}
	}
	PrintOrderList()

}
