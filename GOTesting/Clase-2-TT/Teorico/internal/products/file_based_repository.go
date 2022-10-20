package products

import (
	"tempgo/internal/domain"
	"tempgo/pkg/store"
)

// FileBasedRepository is an implementation of product's Repository interface, based on
// a FileStore structure.
type FileBasedRepository struct {
	store store.Store
}

// NewFileBasedRepository returns a new instance of FileBasedRepository.
func NewFileBasedRepository(store store.Store) *FileBasedRepository {
	return &FileBasedRepository{
		store: store,
	}
}

func (repository *FileBasedRepository) GetAll() (results []domain.Product, err error) {
	err = repository.store.Read(&results)
	return
}

func (repository *FileBasedRepository) Store(productToStore domain.Product) (result domain.Product, err error) {
	err = repository.store.Write(&productToStore)
	if err != nil {
		return
	}

	result = productToStore
	return
}
