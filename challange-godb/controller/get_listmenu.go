package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"database/sql"
)

// func buat nampilin tabel menu dari DB
func GetMenuLoundry() []entity.Menu {
	db := config.ConnectDB()
	defer db.Close()

	SqlCode := "SELECT *FROM list_menu;"

	rows, err := db.Query(SqlCode)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	listMenu := scanMenu(rows)
	return listMenu
}

func scanMenu(rows *sql.Rows) []entity.Menu {
	listMenu := []entity.Menu{}

	var err error
	for rows.Next() {
		menu := entity.Menu{}
		err := rows.Scan(&menu.IdMenu, &menu.JenisLayanan, &menu.Satuan, &menu.Harga)

		if err != nil {
			panic(err)
		}
		listMenu = append(listMenu, menu)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return listMenu
}
