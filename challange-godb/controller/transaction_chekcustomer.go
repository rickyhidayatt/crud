package controller

import (
	"challangge-godb/config"
)

// func ini dibuat untuk validasi jika customer sudah terdaftar atau tidak
func CheckCustomerDB(idOrder int) (bool, error) {

	db := config.ConnectDB()
	defer db.Close()

	var auth bool
	sql := `
		SELECT EXISTS(SELECT 1 FROM transaksi
		JOIN mst_order ON transaksi.id_order = mst_order.id_order 
		WHERE mst_order.id_order = $1 );
	`

	err := db.QueryRow(sql, idOrder).Scan(&auth)
	if err != nil {
		return false, err
	}
	return auth, nil
}
