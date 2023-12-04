package web

import "database/sql"

type RequestBukuFindAll struct {
	Judul string
}

type ResponseBukuFindAll struct {
	IdBuku      int             `json:"id_buku" gorm:"column:id_buku;primaryKey;autoIncrement"`
	Judul       string          `json:"judul"`
	Penulis     sql.NullString  `json:"penulis"`
	Bahasa      sql.NullString  `json:"bahasa"`
	Genre       sql.NullString  `json:"genre"`
	Rating      sql.NullFloat64 `json:"rating"`
	Penerbit    sql.NullString  `json:"penerbit"`
	Description sql.NullString  `json:"description"`
}
