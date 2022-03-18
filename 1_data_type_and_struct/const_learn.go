package main

import (
	"fmt"
	"math"
)

const a string = "abc"
const b = "bbc"
const c, d = "ccc", "ddd"
const (
	Unknow = 0
	Female = 1
	Male   = 2
)

func triangle() {

	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func triangleConst() {

	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func main() {

	fmt.Println(a)
	fmt.Println(d)
	sex := 1
	if Unknow == sex {
		fmt.Println("Unknow")
	} else if Female == sex {
		fmt.Println("Female")
	} else if Male == sex {
		fmt.Println("Male")
	}
}
