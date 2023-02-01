package controller

import (
	"challangge-godb/config"
	"challangge-godb/entity"
	"database/sql"
)

// func buat nampilin tabel customer order dari DB
func GetOrderLoundry() []entity.Order {
	db := config.ConnectDB()
	defer db.Close()

	sql := "SELECT *FROM mst_order ORDER BY id_order;"
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	listOrder := scanOrder(rows)
	return listOrder

}

func scanOrder(rows *sql.Rows) []entity.Order {
	listOrder := []entity.Order{}

	var err error
	for rows.Next() {
		order := entity.Order{}
		err := rows.Scan(&order.IdOrder, &order.NamaCustomer, &order.NamaKasir, &order.NoHP, &order.TotalTagihan)

		if err != nil {
			panic(err)
		}
		listOrder = append(listOrder, order)
	}
	if err != nil {
		panic(err)
	}
	return listOrder
}
