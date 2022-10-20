package products

import (
	"fmt"
	"github.com/dabilaCo/backpack-bcgow6-daniel-abila/GOTesting/Clase-2-TM/Practica/pkg/store"
	
)

// Generamos una estructura
type request struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"` //binding:"required" para indicar que ese campo es requerido
	TypeRequest string  `json:"typeRequest" binding:"required"`
	Amount      int     `json:"amount" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
}

// Generamos una lista para ir guardando los datos que vamos recibiendo
var arrayProducts []request
var lastID int

type Repository interface {
	GetAll() ([]request, error)
	Store(id int, name, typeRequest string, amount int, price float64) (request, error)
	LastID() (int, error)
	Update(id int, name, typeRequest string, amount int, price float64) (request, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, name string, price float64) (request, error)
}

type repository struct{
	db store.Store
} //implementa los metodos de la interfaz

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}	
}

func (r *repository) GetAll() (product[]request, err error) {
	err = r.db.Read(&product)
	if err != nil{
		return nil, err
	}
	return product, nil
}
func (r *repository) Store(id int, name, typeRequest string, amount int, price float64) (request, error) {

	var productSlice []request

	err := r.db.Read(&productSlice)
	if err != nil{
		return request{}, err
	}

	req := request{id, name, typeRequest, amount, price}

	productSlice = append(productSlice, req)
	if err := r.db.Write(productSlice); err != nil{
		return request{}, err
	}	
	return req, nil
}
func (r *repository) LastID() (int, error) {
	var requestSlice []request
	err := r.db.Read(&requestSlice)
	if err != nil{
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
	for i := range arrayProducts {
		if arrayProducts[i].ID == id {
			req.ID = id
			arrayProducts[i] = req
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
	for i := range arrayProducts {
		if arrayProducts[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Producto %d no encontrada", id)
	}
	arrayProducts = append(arrayProducts[:index], arrayProducts[index+1:]...)
	return nil
}
func (r *repository) UpdateNameAndPrice(id int, name string, price float64) (request, error) {
	var requestToReturn request
	updated := false
	for _, requestToUpdate := range arrayProducts {
		if requestToUpdate.ID == id {
			requestToUpdate.Name = name
			requestToUpdate.Price = price
			requestToReturn = requestToUpdate
		}
	}
	if !updated {
		return request{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return requestToReturn, nil
}
