package web

type RequestBukuUpdate struct {
	IdBuku      int     `json:"id_buku"`
	Judul       string  `json:"judul"`
	Penulis     string  `json:"penulis"`
	Bahasa      string  `json:"bahasa"`
	Genre       string  `json:"genre"`
	Rating      float32 `json:"rating"`
	Penerbit    string  `json:"penerbit"`
	Description string  `json:"description"`
}
