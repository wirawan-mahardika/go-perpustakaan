package repo

import (
	"context"
	"perpustakaan/database/entity"
)

type Buku_repo interface {
	Insert(ctx context.Context, buku entity.Buku) (entity.Buku, error)
	FindById(ctx context.Context, id int32) (entity.Buku, error)
	FindAll(ctx context.Context) (entity.Buku, error)
}