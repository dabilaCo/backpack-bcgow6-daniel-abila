package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockedRepository struct{}

func newMockedRepository() Repository {
	return &mockedRepository{}
}

func (m *mockedRepository) GetAllBySeller(id string) ([]Product, error){
	return nil, errors.New("There is an error")
}
func Test_GetAllBySeller(t *testing.T){
	//El repo ya contiene información, no es necesario agregarle la misma
	repo := NewRepository()
	svc := NewService(repo)

	data, err := svc.GetAllBySeller("FEX112AC")
	assert.Nil(t, err)
	assert.Contains(t, data, Product{
		ID: "mock",
		SellerID: "FEX112AC",
		Description: "generic product",
		Price: 123.55,
	})

	repo = newMockedRepository()
	svc = NewService(repo)
	//Esto me va a devolver un error
	data, err = svc.GetAllBySeller("Nothing")
	assert.Nil(t, data)
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("There is an error"))

}
