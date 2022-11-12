package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "p@ssword"
	dbname   = "dbjual"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")

	sqlStatement := "INSERT INTO barang (bar_nama, bar_harga, bar_stok, bar_create, bar_update) " +
		"VALUES ($1, $2, $3, current_timestamp, current_timestamp) RETURNING idbar"
	idbar := 0
	err = db.QueryRow(sqlStatement, "KABEL DATA", 35000, 30).Scan(&idbar)
	if err != nil {
		panic(err)
	}
	fmt.Println("Input Berhasil. ID Barang adalah:", idbar)
}
