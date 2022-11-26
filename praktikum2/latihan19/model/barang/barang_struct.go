package model_barang

import (
	"time"

	_ "github.com/lib/pq"
)

type Barang struct {
	IDBar      string
	Bar_Nama   string
	Bar_Harga  float64
	Bar_Stok   int16
	Bar_Create time.Time
	Bar_Update time.Time
}

type BrgData struct {
	Status string
	Data   []Barang
}

type BrgEnc struct {
	Status string
	Data   string
}
