package main

import "fmt"

func main() {

	//var array1 [5]int是值类型 [5]int，new([5]int)是引用类型 *[5]int.
	var array1 [5]int
	for i := 0; i < len(array1); i++ {
		array1[i] = i
	}

	for index, value := range array1 {
		fmt.Printf("Array at index %d is %d \n", index, value)
	}

	var arr1 = new([5]int)
	arr2 := arr1
	arr2[2] = 100
	arr3 := array1
	fmt.Printf("arr1 at index 2 is: %d, arr2 at index 2 is: %d", arr1[2], arr2[2])
	fmt.Println()
	fmt.Printf("arr1 location in memory: %v, arr2 location in memory: %v", &arr1[1], &arr2[1])
	fmt.Println()
	fmt.Printf("array1 location in memory: %v, arr3 location in memory: %v", &array1[1], &arr3[1])
}
