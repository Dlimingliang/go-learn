package main

import (
	"fmt"
)

func main() {
	s := "{[]}"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	slice := make([]rune, 0)
	for _, v := range []rune(s) {
		if v == '(' || v == '[' || v == '{' {
			slice = append(slice, v)
		} else {
			if len(slice) > 0 &&
				(slice[len(slice)-1] == '(' && v == ')' ||
					slice[len(slice)-1] == '[' && v == ']' ||
					slice[len(slice)-1] == '{' && v == '}') {
				slice = slice[:len(slice)-1]
				continue
			} else {
				return false
			}
		}
	}
	if len(slice) > 0 {
		return false
	}
	return true
}
