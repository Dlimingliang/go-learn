package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {

	//fmt.Println("Please enter you full name:")
	//fmt.Scanln(&firstName, &lastName)
	//fmt.Printf("hi %s %s!\n", firstName, lastName)
	//fmt.Sscanf(input, format, &f, &i, &s)
	//fmt.Println("From the string we read:", f, i, s)111

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
