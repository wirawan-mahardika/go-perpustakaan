package web

type RequestBukuInsert struct {
	IdBuku      int     `json:"id_buku"`
	Judul       string  `json:"judul"`
	Penulis     string  `json:"penulis"`
	Bahasa      string  `json:"bahasa"`
	Genre       string  `json:"genre"`
	Rating      float64 `json:"rating"`
	Penerbit    string  `json:"penerbit"`
	Description string  `json:"description"`
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
}
