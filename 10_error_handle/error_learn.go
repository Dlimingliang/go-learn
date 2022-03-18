package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
	fmt.Println()
	if _, err := sqrt(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math - square root of negative number")
	}
	return f, nil
}
