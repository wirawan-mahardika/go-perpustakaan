package service

import (
	"context"
	exception "perpustakaan/error"
	"perpustakaan/model/domain"
	"perpustakaan/model/web"

	"gorm.io/gorm"
)

type BukuService interface {
	FindAllOrSearch(ctx context.Context, request web.RequestBukuFindAll) []web.ResponseBukuFindAll
	FindById(ctx context.Context, request web.RequestBukuFindById) web.ResponseBukuFindById
	Insert(ctx context.Context, request web.RequestBukuInsert)
	Update(ctx context.Context, request web.RequestBukuUpdate)
}

func NewBukuService(db *gorm.DB) BukuService {
	return &bukuServiceImpl{db: db}
}

type bukuServiceImpl struct {
	db *gorm.DB
}

func (service *bukuServiceImpl) FindAllOrSearch(ctx context.Context, request web.RequestBukuFindAll) []web.ResponseBukuFindAll {
	responseBuku := []web.ResponseBukuFindAll{}
	var err error
	if request.Judul != "" {
		err = service.db.Model(&domain.Buku{}).Select("id_buku", "judul", "penerbit", "penulis", "genre", "rating").Where("judul like ?", "%"+request.Judul+"%").Find(&responseBuku).Error
	} else {
		err = service.db.Model(&domain.Buku{}).Select("id_buku", "judul", "penerbit", "penulis", "genre", "rating").Find(&responseBuku).Error
	}
	if err != nil {
		panic(err)
	}

	if len(responseBuku) < 1 {
		panic(exception.NotFound{Message: "Buku tidak ditemukan"})
	}
	return responseBuku
}

func (service *bukuServiceImpl) FindById(ctx context.Context, request web.RequestBukuFindById) web.ResponseBukuFindById {
	responseBuku := web.ResponseBukuFindById{}
	result := service.db.Model(&domain.Buku{}).Take(&responseBuku, "id_buku = ?", request.Id)

	if result.RowsAffected == 0 {
		panic(exception.NotFound{Message: "Buku tidak ditemukan"})
	}

	return responseBuku
}

func (service *bukuServiceImpl) Insert(ctx context.Context, request web.RequestBukuInsert) {
	err := service.db.Omit("id_buku", "created_at", "updated_at").Model(&domain.Buku{}).Create(&request).Error
	if err != nil {
		panic(err)
	}
}

func (service *bukuServiceImpl) Update(ctx context.Context, request web.RequestBukuUpdate) {
	err := service.db.Model(&domain.Buku{}).Omit("id_buku").Where("id_buku = ?", request.IdBuku).Updates(&request).Error
	if err != nil {
		panic(err)
	}
}
