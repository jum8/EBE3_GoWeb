package product

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/jum8/EBE3_GoWeb.git/internal/domain"
)

var (
	ErrPrepareStmt    = errors.New("error preparing statement")
	ErrExecStmt       = errors.New("error executing statement")
	ErrRowScan        = errors.New("error scanning row")
	ErrQuery          = errors.New("error querying")
	ErrLastInsertedId = errors.New("error getting last inserted id")
)

type repositorySql struct {
	storage *sql.DB
}

func NewSqlRespository(db *sql.DB) Respository {
	return &repositorySql{storage: db}

}

// GetAll implements Respository.
func (r *repositorySql) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product

	rows, err := r.storage.Query(selectAllProducts)
	if err != nil {
		return nil, ErrQuery
	}

	for rows.Next() {
		var currentProduct domain.Product
		err := rows.Scan(&currentProduct.Id, &currentProduct.Name, &currentProduct.Quantity, &currentProduct.CodeValue,
			&currentProduct.IsPublished, &currentProduct.Expiration, &currentProduct.Price)

		if err != nil {
			return nil, ErrRowScan
		}

		products = append(products, currentProduct)

	}

	return products, nil
}

// GetById implements Respository.
func (r *repositorySql) GetById(ctx context.Context, id string) (domain.Product, error) {
	var product domain.Product

	row := r.storage.QueryRow(selectProductById, id)
	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)

	if err != nil {
		return domain.Product{}, ErrRowScan
	}

	return product, nil
}

// Save implements Respository.
func (r *repositorySql) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	productSaved := copyProduct(product)

	stmt, err := r.storage.Prepare(insertProduct)
	if err != nil {
		return domain.Product{}, ErrPrepareStmt
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
	)

	if err != nil {
		return domain.Product{}, ErrExecStmt
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, ErrLastInsertedId
	}

	productSaved.Id = strconv.Itoa(int(id))

	return productSaved, nil
}

// Update implements Respository.
func (r *repositorySql) Update(ctx context.Context, product domain.Product, id string) (domain.Product, error) {
	productUpdated := copyProduct(product)

	stmt, err := r.storage.Prepare(updateProduct)
	if err != nil {
		return domain.Product{}, ErrPrepareStmt
	}

	_, err = stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, id)

	if err != nil {
		return domain.Product{}, ErrExecStmt
	}

	productUpdated.Id = id

	return productUpdated, nil
}

// Delete implements Respository.
func (r *repositorySql) Delete(ctx context.Context, id string) error {
	stmt, err := r.storage.Prepare(deleteProduct)
	if err != nil {
		return ErrPrepareStmt
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return ErrExecStmt
	}

	return nil
}

// copies one product fields into another one
func copyProduct(source domain.Product) domain.Product {
	var destination domain.Product

	destination.Id = source.Id
	destination.Name = source.Name
	destination.Quantity = source.Quantity
	destination.CodeValue = source.CodeValue
	destination.IsPublished = source.IsPublished
	destination.Expiration = source.Expiration
	destination.Price = source.Price

	return destination
}
