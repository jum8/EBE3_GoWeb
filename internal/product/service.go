package product

import (
	"context"
	"log"

	"github.com/jum8/EBE3_GoWeb.git/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetById(ctx context.Context, id string) (domain.Product, error)
	Save(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id string) (domain.Product, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Respository
}

func NewProductService(repo Respository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	productsList, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[ProductsService][GetAll] error getting all products", err)
		return []domain.Product{}, err
	}
	return productsList, nil
}

func (s *service) GetById(ctx context.Context, id string) (domain.Product, error) {
	productFound, err := s.repository.GetById(ctx, id)
	if err != nil {
		log.Println("[ProductService][GetById] error getting product by id	", err)
		return domain.Product{}, err
	}
	return productFound, nil
}

func (s *service) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	product, err := s.repository.Save(ctx, product)
	if err != nil {
		log.Println("[ProductService][Save] error saving a product", err)
		return domain.Product{}, err
	}
	return product, nil
}


// Update implements Service.
func (s *service) Update(ctx context.Context, product domain.Product, id string) (domain.Product, error) {
	result, err := s.repository.Update(ctx, product, id)
	if err != nil {
		log.Println("[ProductsService][Update] error updating product by ID", err)
		return domain.Product{}, err
	}
	return result, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[ProductService][Delete] error deleting product", err)
		return err
	}
	return nil
}
