package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock       []request
	readWasCalled  bool
	writeWasCalled bool
	errWrite       string
	errRead        string
}

func (m *MockStorage) Read(data interface{}) error {
	m.readWasCalled = true
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]request)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	m.writeWasCalled = true
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]request)
	m.dataMock = a
	return nil
}

// Ejercicio 1
func TestServiceIntegrationUpdate(t *testing.T) {
	//Arrange
	product := request{
		ID:          1,
		Name:        "Before Update",
		TypeRequest: "7287828727",
		Amount:      26,
		Price:       2000,
	}
	updated := request{
		ID:          1,
		Name:        "After Update",
		TypeRequest: "7287828727",
		Amount:      4,
		Price:       5000,
	}

	db := []request{
		product,
	}

	mockedStorage := MockStorage{
		dataMock: db,
	}
	rp := NewRepository(&mockedStorage)
	svc := NewService(rp)

	//Act
	result, err := svc.Update(
		updated.ID,
		updated.Name,
		updated.TypeRequest,
		updated.Amount,
		updated.Price)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, result, updated)
	assert.True(t, mockedStorage.readWasCalled)

}

// Ejercicio 2
func TestServiceIntegrationDelete(t *testing.T) {
	//Arrange
	product1 := request{
		ID:          1,
		Name:        "Monitor Gamer",
		TypeRequest: "CVH-432",
		Amount:      26,
		Price:       4500,
	}
	product2 := request{
		ID:          2,
		Name:        "Teclado Gamer",
		TypeRequest: "CVH-444",
		Amount:      26,
		Price:       1500,
	}

	db := []request{product1, product2}

	mockedStorage := MockStorage{dataMock: db}
	rp := NewRepository(&mockedStorage)
	svc := NewService(rp)

	//Act
	err := svc.Delete(2)
	assert.Nil(t, err)
	assert.True(t, mockedStorage.readWasCalled)
	assert.True(t, mockedStorage.writeWasCalled)
	assert.Equal(t, mockedStorage.dataMock, []request{product1})

	err = svc.Delete(2)
	assert.NotNil(t, err)
	assert.True(t, mockedStorage.readWasCalled)
	assert.True(t, mockedStorage.writeWasCalled)
	assert.Equal(t, mockedStorage.dataMock, []request{product1})
}
