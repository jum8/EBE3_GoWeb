package product

import (
	"context"
	"errors"
	"time"

	"github.com/jum8/EBE3_GoWeb.git/internal/domain"
)

var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
)


type repository struct {
	products []domain.Product
}


func NewMemoryRespository() Respository {
	return &repository{
		products: []domain.Product{
			{
				Id:          "1",
				Name:        "Coco Cola",
				Quantity:    10,
				CodeValue:   "123456789",
				IsPublished: true,
				Expiration:  time.Now(),
				Price:       10.5,
			},
			{
				Id:          "2",
				Name:        "Pepsito",
				Quantity:    10,
				CodeValue:   "123456789",
				IsPublished: true,
				Expiration:  time.Now(),
				Price:       8.5,
			},
			{
				Id:          "3",
				Name:        "Fantastica",
				Quantity:    10,
				CodeValue:   "123456789",
				IsPublished: true,
				Expiration:  time.Now(),
				Price:       5.5,
			},
		},
	}
	
}

// Get all products
func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	if len(r.products) < 1 {
		return []domain.Product{}, ErrEmpty
	}
	return r.products, nil
}

// Save a product
func (r *repository) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	r.products = append(r.products, product)
	return product, nil
}

// Update implements Respository.
func (r *repository) Update(ctx context.Context, product domain.Product, id string) (domain.Product, error) {
	var result domain.Product
	for index, currentProduct := range r.products {
		if currentProduct.Id == id {
			product.Id = id
			r.products[index] = product
			result = r.products[index]
			break
		}
	}

	if result.Id == "" {
		return domain.Product{}, ErrNotFound
	}
	
	return result, nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	var result domain.Product
	for index, product := range r.products {
		if product.Id == id {
			result = r.products[index]
			r.products = append(r.products[:index], r.products[index+1:]...)
			break
		}
	}

	if result.Id == "" {
		return ErrNotFound
	}
	return nil
}

// Get by id
func (r *repository) GetById(ctx context.Context, id string) (domain.Product, error) {
	var productFound domain.Product
	for _, product := range r.products {
		if product.Id == id {
			productFound = product
			break
		}
	}

	if productFound.Id == "" {
		return domain.Product{}, ErrNotFound
	}

	return productFound, nil
}
