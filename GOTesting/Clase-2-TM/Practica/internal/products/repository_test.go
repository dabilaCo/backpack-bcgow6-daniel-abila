package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubReadEngine struct {
	mockedData    []request
	readWasCalled bool
}

func (g *StubReadEngine) Read(data interface{}) error {

	g.readWasCalled = true
	a, ok := data.(*[]request)
	if !ok {
		return errors.New("it failed!")
	}
	*a = g.mockedData
	return nil
}

func (g *StubReadEngine) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	expectedProducts := []request{
		{
			ID:          1,
			Name:        "Mouse",
			TypeRequest: "7287828727",
			Amount:      26,
			Price:       2,
		},
		{
			ID:          2,
			Name:        "Teclado",
			TypeRequest: "ACH-VC5",
			Amount:      26,
			Price:       1500,
		},
	}
	db := &StubReadEngine{mockedData: expectedProducts}
	r := NewRepository(db)

	out, err := r.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedProducts, out)
}

func TestUpdateNameAndPrice(t *testing.T) {
	productos := []request{
		{

			ID:          1,
			Name:        "Before Update",
			TypeRequest: "7287828727",
			Amount:      26,
			Price:       2000,
		},
		{
			ID:          2,
			Name:        "Teclado",
			TypeRequest: "ACH-VC5",
			Amount:      26,
			Price:       1500,
		},
	}

	expectedProduct := productos[0]
	updatedName := "After Update"
	updatedPrice := 2500
	expectedProduct.Name = updatedName
	expectedProduct.Price = float64(updatedPrice)
	db := &StubReadEngine{mockedData: productos, readWasCalled: false}
	r := NewRepository(db)

	out, err := r.UpdateNameAndPrice(1, updatedName, float64(updatedPrice))

	assert.Nil(t, err)
	assert.Equal(t, expectedProduct, out)
	assert.True(t, db.readWasCalled)

}
