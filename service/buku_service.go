package service

import (
	"context"
	"database/sql"
	"perpustakaan/model/web"
	"perpustakaan/repository"
)

type BukuService interface {
	FindById(ctx context.Context, request web.RequestBukuFindById) (response web.ResponseBuku)
}

type BukuServiceImpl struct {
	DB   *sql.DB
	Repo repository.BukuRepository
}

func NewBukuService(db *sql.DB, repo repository.BukuRepository) BukuService {
	return &BukuServiceImpl{DB: db, Repo: repo}
}

func (service *BukuServiceImpl) FindById(ctx context.Context, request web.RequestBukuFindById) (response web.ResponseBuku) {
	db := service.DB

	responseBuku := service.Repo.FindById(ctx, db, request.Id)
	response = web.ResponseBuku{
		Id_buku:       responseBuku.Id_buku,
		Judul:         responseBuku.Judul,
		Penulis:       responseBuku.Penulis,
		Bahasa:        responseBuku.Bahasa,
		Genre:         responseBuku.Genre,
		Penerbit:      responseBuku.Penerbit,
		TanggalTerbit: responseBuku.TanggalTerbit,
		Description:   responseBuku.Description,
	}

	return response
}
