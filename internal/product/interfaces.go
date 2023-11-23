package product

import (
	"context"

	"github.com/jum8/EBE3_GoWeb.git/internal/domain"
)

type Respository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id int) (domain.Product, error)
	Delete(ctx context.Context, id int) error
}