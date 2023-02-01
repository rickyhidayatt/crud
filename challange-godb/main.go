package main

import (
	"challangge-godb/fitur"
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("|", strings.Repeat(" ", 15), "ENIGMA LOUNDRY", strings.Repeat(" ", 15), "|")
	fmt.Println(strings.Repeat("=", 50))

	fmt.Println("Hallo selamat datang di aplikasi enigma loundry !")
	fitur.MenuApps()

}
