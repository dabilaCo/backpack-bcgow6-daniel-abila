package products

import (
	"fmt"

	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-3-TT/Practica/pkg/store"
)

// Generamos una estructura
type request struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"` //binding:"required" para indicar que ese campo es requerido
	TypeRequest string  `json:"typeRequest" binding:"required"`
	Amount      int     `json:"amount" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}


type Repository interface {
	GetAll() ([]request, error)
	Store(id int, name, typeRequest string, amount int, price float64) (request, error)
	LastID() (int, error)
	Update(id int, name, typeRequest string, amount int, price float64) (request, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, name string, price float64) (request, error)
}

type repository struct {
	db store.Store
} //implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() (product []request, err error) {
	err = r.db.Read(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (r *repository) Store(id int, name, typeRequest string, amount int, price float64) (request, error) {

	var productSlice []request

	err := r.db.Read(&productSlice)
	if err != nil {
		return request{}, err
	}

	req := request{id, name, typeRequest, amount, price}

	productSlice = append(productSlice, req)
	if err := r.db.Write(productSlice); err != nil {
		return request{}, err
	}
	return req, nil
}
func (r *repository) LastID() (int, error) {
	var requestSlice []request
	err := r.db.Read(&requestSlice)
	if err != nil {
		return 0, err
	}
	if len(requestSlice) == 0 {
		return 0, nil
	}
	return requestSlice[len(requestSlice)-1].ID, nil
}
func (r *repository) Update(id int, name, typeRequest string, amount int, price float64) (request, error) {
	req := request{Name: name, TypeRequest: typeRequest, Amount: amount, Price: price}
	updated := false
	var ps []request
	if err := r.db.Read(&ps); err != nil {
		return request{}, err
	}
	for i := range ps {
		if ps[i].ID == id {
			req.ID = id
			ps[i] = req
			updated = true
		}
	}
	if !updated {
		return request{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return req, nil
}


func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	var ps []request
	if err := r.db.Read(&ps); err != nil {
		return err
	}
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Producto %d no encontrada", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	if r.db.Write(ps) != nil {
		return fmt.Errorf("error while deleting product %d", id)
	}
	return nil
}
func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (result request, err error) {
	//obtener todos los productos del store
	var productsInStore []request
	err = r.db.Read(&productsInStore)
	if err != nil {
		return
	}

	//recorremos cada elemento en busqueda de el producto a actualizar
	var updated bool
	for i := range productsInStore {
		if productsInStore[i].ID == id {
			productsInStore[i].Name = name
			productsInStore[i].Price = price
			updated = true
			result = productsInStore[i]
			break
		}
	}

	if !updated {
		err = fmt.Errorf("Product not found")
		return
	}

	err = r.db.Write(&result)
	if err != nil {
		result = request{}
		return
	}
	return
}
