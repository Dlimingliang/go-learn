package main

import (
	"fmt"
	"time"
)

func main() {

	a := 1
	fmt.Println(a)
	changeValue(a)
	fmt.Println(a)
	fmt.Println("three value mutiply result is:", mutiplyThreeValue(1, 2, 3))

	i1, _, f1 := threeValues()
	fmt.Println("i1", i1)
	fmt.Println("f1", f1)

	n := 0
	mutiply(5, 10, &n)
	fmt.Println(n)

	fmt.Printf("The minimum is: %d\n", min(1, 2, 3, 0))
	slice := []int{1, 2, 3, 4}
	fmt.Printf("The minimum is: %d\n", min(slice...))

	fmt.Println(getX2AndX3(2))

	callBack(1, add)

	fpluus := func(x, y int) int {
		return x + y
	}
	fmt.Println(fpluus(3, 4))
	func(x, y int) int {
		return x + y
	}(2, 5)

	fmt.Println(f())

	f := adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	fmt.Println()

	start := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func changeValue(a int) {
	a = 2
	fmt.Println(a)
}

func mutiplyThreeValue(a, b, c int) int {
	return a * b * c
}

func getX2AndX3(input int) (x2, x3 int) {
	defer fmt.Printf("getX2AndX3Parameter: %d, result x2: %d, result x3: %d", input, x2, x3)
	defer fmt.Println()
	defer func() {
		fmt.Printf("getX2AndX3Parameter: %d, result x2: %d, result x3: %d", input, x2, x3)
	}()
	x2 = 2 * input
	x3 = 3 * input
	return
}

func threeValues() (int, int, float32) {
	return 5, 6, 7.5
}

func mutiply(a, b int, reply *int) {
	*reply = a * b
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func sumArgs(values ...int) int {
	sum := 0
	for v := range values {
		sum += v
	}
	return sum
}

func add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callBack(input int, f func(int, int)) {
	f(input, 2)
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}
