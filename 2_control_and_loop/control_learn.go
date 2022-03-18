package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}

	var orig string = "11"
	var newS string

	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	var num1 int = 100
	switch num1 {
	case 98, 99:
		fmt.Println("1")
	case 100:
		fmt.Println("2")
	default:
		fmt.Println("0")
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	fmt.Println(season(time.Now().Day()))
}

func season(num int) (str string) {

	switch num {
	case 1, 2, 3:
		return "春"
	case 4, 5, 6:
		return "夏"
	case 7, 8, 9:
		return "秋"
	case 10, 11, 12:
		return "冬"
	}
	return ""
}
