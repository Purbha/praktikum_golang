package model_barang

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	config_db "github.com/purbha/latihan/praktikum2/latihan16/config"
)

func getpsqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config_db.DBHost, config_db.DBPort, config_db.DBUser, config_db.DBPass, config_db.DBName)
}

func AmbilBarangs(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlInfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	var barangs []Barang
	//var brgdata BrgData
	var brgenc BrgEnc
	result, err := db.Query("SELECT * FROM barang")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var barang Barang
		err := result.Scan(&barang.IDBar, &barang.Bar_Nama, &barang.Bar_Harga, &barang.Bar_Stok,
			&barang.Bar_Create, &barang.Bar_Update)
		if err != nil {
			panic(err.Error())
		}
		barangs = append(barangs, barang)
	}
	/*
		encText, err := Encrypt(barangs, MySecret)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
		}
	*/
	/*
		brgdata.Status = "OK"
		brgdata.Data = barangs
	*/
	e, err := json.Marshal(barangs)
	if err != nil {
		panic(err.Error())
	}
	encText, err := Encrypt(string(e), MySecret)
	if err != nil {
		panic(err.Error())
	}
	brgenc.Status = "OK"
	brgenc.Data = encText
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(barangs)
	json.NewEncoder(w).Encode(brgenc)
}

func AmbilBarang(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", getpsqlInfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Koneksi Sukses")
	params := mux.Vars(r)
	idbar := params["id"]
	var barang Barang
	sqlStatement := "SELECT * FROM barang WHERE idbar = " + idbar
	result, err := db.Query(sqlStatement)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&barang.IDBar, &barang.Bar_Nama, &barang.Bar_Harga, &barang.Bar_Stok,
			&barang.Bar_Create, &barang.Bar_Update)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(barang)
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b35"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}
