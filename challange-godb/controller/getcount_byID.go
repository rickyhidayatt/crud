package controller

import (
	"challangge-godb/config"
)

func GetIDCount() (int, error) {
	db := config.ConnectDB()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM list_menu").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
