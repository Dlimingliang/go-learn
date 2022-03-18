package main

import "fmt"

func main() {

	var mapList map[string]int
	var mapAssigned map[string]int

	mapList = map[string]int{"one": 1, "two": 2}
	mapCreate := make(map[string]float32)
	mapAssigned = mapList

	mapCreate["key1"] = 4.5
	mapCreate["key2"] = 3.1415926
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreate["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapList["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapList["ten"])

	_, ok := mapCreate["key1"]
	fmt.Println("map wether contains key1", ok)
	delete(mapCreate, "key1")
	_, ok = mapCreate["key1"]
	fmt.Println("map wether contains key1", ok)

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}
