package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	str := "Go is beautiful language"
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is: %c \n", i, str[i])
	}

	num := 4
	for num > 0 {
		fmt.Println(num)
		num -= 1
	}

	for i, i2 := range str {
		fmt.Printf("Character on position %d is: %c \n", i, i2)
	}

	for i := 0; i < 5; i++ {
		if i == 0 {
			continue
		}
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
}
