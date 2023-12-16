package web

type RequestBukuFindAll struct {
	Judul string
}

type ResponseBukuFindAll struct {
	IdBuku      int      `json:"id_buku" gorm:"column:id_buku;primaryKey;autoIncrement"`
	Judul       string   `json:"judul"`
	Penulis     *string  `json:"penulis"`
	Bahasa      *string  `json:"bahasa"`
	Genre       *string  `json:"genre"`
	Rating      *float32 `json:"rating"`
	Penerbit    *string  `json:"penerbit"`
	Description *string  `json:"description"`
}
