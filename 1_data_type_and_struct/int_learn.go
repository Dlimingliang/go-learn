package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//var i1 int = 1
	var i1 = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5

	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(i3)
	fmt.Println(i4)
	fmt.Println(i5)

	fmt.Println(unsafe.Sizeof(i1))
	fmt.Println(unsafe.Sizeof(i2))
	fmt.Println(unsafe.Sizeof(i3))
	fmt.Println(unsafe.Sizeof(i4))
	fmt.Println(unsafe.Sizeof(i5))
}
