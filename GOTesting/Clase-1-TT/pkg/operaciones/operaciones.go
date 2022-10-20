package operaciones

import "fmt"

// Add: Add two numbers
func Add(a int, b int) (int, error) {
	if a == 0 {
		return 0, fmt.Errorf("a no puede ser:%d", a)
	}
	return a + b, nil
}

// Sub: Substract two numbers
func Sub(a int, b int) int {
	return a - b
}

func add() {}
