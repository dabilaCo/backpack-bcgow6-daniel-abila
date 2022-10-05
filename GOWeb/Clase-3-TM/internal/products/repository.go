package products

import (
	"fmt"
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

type repository struct{} //implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]request, error) {
	return arrayProducts, nil
}
func (r *repository) Store(id int, name, typeRequest string, amount int, price float64) (request, error) {
	req := request{id, name, typeRequest, amount, price}
	arrayProducts = append(arrayProducts, req)
	lastID = req.ID
	return req, nil
}
func (r *repository) LastID() (int, error) {
	return lastID, nil
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
