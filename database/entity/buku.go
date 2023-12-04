package entity

import "time"

type Buku struct {
	Id_buku      int32    `json:"id_buku"`
	Nama         string `json:"nama"`
	Penerbit     string `json:"penerbit"`
	Tahun_terbit time.Time `json:"tahun_terbit"`
	Penulis string `json:"penulis"`
	Updated_at time.Time `json:"updated_at"`
}