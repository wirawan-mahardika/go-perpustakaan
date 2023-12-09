package entity

import (
	"database/sql"
)

type Buku struct {
	Id_buku       int            `json:"id_buku"`
	Judul         string         `json:"judul"`
	Penulis       sql.NullString `json:"penulis"`
	Bahasa        sql.NullString `json:"bahasa"`
	Genre         sql.NullString `json:"genre"`
	Penerbit      sql.NullString `json:"penerbit"`
	TanggalTerbit sql.NullTime   `json:"tanggal_terbit"`
	Description   sql.NullString `json:"description"`
}
