package repository

import (
	"context"
	"database/sql"
	"perpustakaan/helper"
	"perpustakaan/model/entity"
)

type BukuRepository interface {
	// Insert(ctx context.Context, db *sql.DB, buku entity.Buku) (message string)
	FindById(ctx context.Context, db *sql.DB, id int) (buku entity.Buku)
}

type BukuRepositoryImpl struct{}

func NewBukuRepository() BukuRepository {
	return &BukuRepositoryImpl{}
}

func (repo *BukuRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id int) (buku entity.Buku) {
	script := "SELECT id_buku, judul, penulis, bahasa, genre, penerbit, tanggal_terbit, description FROM buku WHERE id_buku = ?;"
	rows, err := db.QueryContext(ctx, script, id)
	helper.PanicIfError(err)

	if rows.Next() {
		err := rows.Scan(&buku.Id_buku, &buku.Judul, &buku.Penulis, &buku.Bahasa, &buku.Genre, &buku.Penerbit, &buku.TanggalTerbit, &buku.Description)
		helper.PanicIfError(err)
	} else {
		panic("Buku tidak ditemukan")
	}
	return buku
}

// func (repo *BukuRepositoryImpl) Insert(ctx context.Context, buku entity.Buku) (message string) {
// 	script := "INSERT buku(judul, penulis, bahasa, genre, penerbit, tahun_terbit, description)INTO VALUES (?,?,?,?,?,?,?);"
// args := []any{
// 	buku.Id_buku,
// 	buku.Judul,
// 	buku.Penulis,
// 	buku.Bahasa,
// 	buku.Genre,
// 	buku.Penerbit,
// 	buku.TanggalTerbit,
// 	buku.Description,
// }
// }
