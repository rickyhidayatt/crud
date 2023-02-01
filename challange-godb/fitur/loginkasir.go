package fitur

import "fmt"

func Login(password string) (bool, error) {
	var masuk bool

	for !masuk {
		// check password
		if password == "Rickyganteng" {
			masuk = true
		} else {
			fmt.Println("Password salah, silakan coba lagi.")
			var input string
			fmt.Scan(&input)
			password = input
		}
	}
	return masuk, nil
}
