package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader(
		"Hello World",
	)

	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		panic(err)
	}
}
