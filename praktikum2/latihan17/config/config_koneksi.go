package config

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB
var Err error

func getpsqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPass, DBName)
}

func Buka_koneksi() {
	Db, Err = sql.Open("postgres", getpsqlInfo())
	if Err != nil {
		panic(Err)
	}
	fmt.Println("Koneksi Sukses")
}

func Tutup_koneksi() {
	defer Db.Close()
}
