package main

import (
	"errors"
	"fmt"
)

type MyCustomError struct {
	StatusCode int
	Message    string
}

func (err *MyCustomError) Error() string {
	return fmt.Sprintf("%s (%d)", err.Message, err.StatusCode)
}

func main() {
	err1 := &MyCustomError{
		StatusCode: 502,
		Message:    "Soy el error #1, soy distinto al #2",
	}

	err2 := &MyCustomError{
		StatusCode: 406,
		Message:    "Soy el error #2, soy distinto al error #1",
	}

	bothErrorsAreEqual := errors.As(err1, err2)

	fmt.Println(bothErrorsAreEqual)
}
