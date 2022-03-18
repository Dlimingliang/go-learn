package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a float32
	var b float64
	fmt.Println(a)
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println("Hello, world")
}
