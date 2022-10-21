package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubReadEngine struct{}

func (g StubReadEngine) Read(data interface{}) error {
	products := []request{
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
	otro := data.(*[]request)
	*otro = append(*otro, products...)
	return nil
}

func (g StubReadEngine) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	db := StubReadEngine{}
	repository := NewRepository(db)

	prod, err := repository.GetAll()
	if err != nil {
		fmt.Println("Se rompió")
	}

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
	assert.Equal(t, expectedProducts, prod)
}

type MockRead struct {
	readCheck bool
	productos []request
}

func (mk *MockRead) Read(data interface{}) error {
	mk.readCheck = true
	castData := data.(*[]request)
	*castData = mk.productos
	return nil
}

func (mk *MockRead) Write(data interface{}) error {
	castData := data.(*[]request)
	mk.productos = *castData
	return nil
}

func TestUpdateName(t *testing.T) {
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

	db := MockRead{productos: productos}
	repository := NewRepository(&db)

	expectedProduct := request{
		ID:          1,
		Name:        "After Update",
		TypeRequest: "7287828727",
		Amount:      26,
		Price:       5000,
	}

	newT, err := repository.UpdateNameAndPrice(1, "After Update", 5000)
	assert.Nil(t, err)
	assert.True(t, db.readCheck)
	assert.Equal(t, expectedProduct, newT)
}
