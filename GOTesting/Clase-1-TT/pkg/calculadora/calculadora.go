package calculadora

import "errors"

// Resta
func Resta(a int, b int) (int, error) {
	return a - b, nil
}

func Dividir(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("Are you f***in' crazy!?")
	}

	return num1 / num2, nil
}

func resta() {}
