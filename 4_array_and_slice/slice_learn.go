package main

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	fmt.Println("s1 :", s1)
	s2 := s1[3:5]
	fmt.Println("s2 :", s2)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println(s3, s4, s5)
	fmt.Println("arr :", arr)

	var arr1 [6]int
	var slice1 = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of arr1 is %d\n", len(arr1))       //6
	fmt.Printf("The length of slice1 is %d\n", len(slice1))   //3
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) //4

	slice1 = slice1[0:cap(slice1)]
	for i, i2 := range slice1 {
		fmt.Printf("Slice at %d is %d\n", i, i2)
	}

	for i := range slice1 {
		fmt.Printf("Slice index is %d\n", i)
	}

	for _, i2 := range slice1 {
		fmt.Printf("Slice value is %d\n", i2)
	}

	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	fmt.Println(sum([]int{0, 1, 2, 3, 4}))

	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	sl3 = append(sl3, slFrom...)
	fmt.Println(sl3)

	slice := make([]int, 0, 5)
	slicePoint := &slice
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))
	fmt.Println("index0的值: ", slice[0])
	//删除元素并且保留容量
	slice = append(slice[:0], slice[1:]...)
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))
	slice = slice[1:]
	fmt.Println("当前长度: ", len(*slicePoint))
	fmt.Println("当前容量: ", cap(slice))
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
