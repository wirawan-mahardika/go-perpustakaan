package domain

import (
	"database/sql"

	"gorm.io/gorm"
)

type Buku struct {
	IdBuku      int             `json:"id_buku" gorm:"column:id_buku;primaryKey;autoIncrement"`
	Judul       string          `json:"judul"`
	Penulis     sql.NullString  `json:"penulis"`
	Bahasa      sql.NullString  `json:"bahasa"`
	Genre       sql.NullString  `json:"genre"`
	Rating      sql.NullFloat64 `json:"rating"`
	Penerbit    sql.NullString  `json:"penerbit"`
	Description sql.NullString  `json:"description"`
	CreatedAt   sql.NullTime    `json:"created_at" gorm:"autoCreateTime;"`
	UpdatedAt   sql.NullTime    `json:"updated_at" gorm:"autoCreateTime;autoUpdateTime;"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" gorm:"autoUpdateTime;->:false"`
}
