package products

import "tempgo/internal/domain"

// Repository interface for products repositories implementations.
type Repository interface {
	// GetAll returns all products stored in the repository.
	GetAll() ([]domain.Product, error)

	// Store a product in the given storage.
	Store(productToStore domain.Product) (result domain.Product, err error)
}
