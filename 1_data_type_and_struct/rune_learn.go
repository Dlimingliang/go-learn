package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱英雄联盟!"
	fmt.Println(len(s))
	fmt.Printf("%x\n", []byte(s))
	for i, v := range []byte(s) {
		fmt.Printf("(%d %x) ", i, v)
	}
	fmt.Println()
	for i, v := range s { // v is a rune
		fmt.Printf("(%d %x) ", i, v)
	}
	fmt.Println()

	fmt.Println("Rune count", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	for i, v := range []rune(s) {
		fmt.Printf("(%d %c) ", i, v)
	}

}
