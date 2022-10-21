package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num1, num2 := 100, 2
	resultadoEsperado := 50

	resultado, err := Dividir(num1, num2)
	assert.Equal(t, resultado, resultadoEsperado, "Ajaja@")
	assert.Nil(t, err, "Fiuuu!")

	num2 = 0
	resultadoEsperado = 0

	resultado, err = Dividir(num1, num2)
	assert.Equal(t, resultado, resultadoEsperado, "chan?")
	assert.NotNil(t, err, "Roger that!")
}

func TestRestar(t *testing.T) {
	// Arrange
	num1 := 10
	num2 := 5
	esperado := 5
	resta()
	// Act
	resultado, err := Resta(num1, num2)
	// Assert
	assert.Equal(t, esperado, resultado, "El numero resultado: %d, es distinto del esperado: %d ", resultado, esperado)
	assert.Nil(t, err)

}
