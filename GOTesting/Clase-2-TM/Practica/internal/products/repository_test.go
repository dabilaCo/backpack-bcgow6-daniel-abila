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

func (mk *MockRead) Read(data interface{}) error{
	productos := []request{
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
	mk.productos = productos
	productData := data.(*[]request)
	*productData = append(*productData, productos...)
	mk.readCheck = true	
	return nil
}

func (mk MockRead) Write(data interface{}) error{
	productDate := data.([]request)
	mk.productos = productDate
	return nil
}

func TestUpdateName(t *testing.T){
	db := MockRead{}
	repository := NewRepository(&db)

	expectedProduct := request{

			ID:          1,
			Name:        "Mouse",
			TypeRequest: "7287828727",
			Amount:      26,
			Price:       2000,
	}

	newT, err := repository.UpdateNameAndPrice(1, "Mouse", 2000)
	if err != nil{
		fmt.Print("algo")
	}
	assert.True(t, db.readCheck)
	assert.Equal(t, expectedProduct, newT)
}