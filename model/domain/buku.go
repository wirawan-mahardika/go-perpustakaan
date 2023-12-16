package domain

import (
	"time"
)

type Buku struct {
	IdBuku      int       `gorm:"column:id_buku;primaryKey;autoIncrement;"`
	Judul       string    `gorm:"column:judul;default:null"`
	Penulis     string    `gorm:"column:penulis;default:null"`
	Bahasa      string    `gorm:"column:bahasa;default:null"`
	Genre       string    `gorm:"column:genre;default:null"`
	Rating      string    `gorm:"column:rating;default:null"`
	Penerbit    string    `gorm:"column:penerbit;default:null"`
	Description string    `gorm:"column:description;default:null"`
	CreatedAt   time.Time `gorm:"autoCreateTime;"`
	UpdatedAt   time.Time `gorm:"autoCreateTime;autoUpdateTime;"`
}
