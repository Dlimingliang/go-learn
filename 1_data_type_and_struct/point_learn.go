package main

import "fmt"

func main() {
	var pointValue int = 5
	fmt.Printf("value is %d, it`s location in memory is %p", pointValue, &pointValue)
	fmt.Println()
	var point *int = &pointValue
	fmt.Printf("The value at memory location %p is %d", point, *point)
}
