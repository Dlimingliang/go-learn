package main

import (
	"fmt"
	"os"
	"runtime"
)

var variebleA int = 14
var variebleB = 15
var str = "Hello World"
var aa = "G"
var bbb string

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func main() {

	var (
		cc = 1
		dd = 2
	)
	fmt.Println(cc, dd)
	one := 1
	fmt.Println(one)
	fmt.Println(HOME)
	var goos = runtime.GOOS
	fmt.Printf("The operating system is %s\n", goos)
	fmt.Printf("Path is %s\n", os.Getenv("PATH"))

	i := 7
	j := i
	i = 8
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(&i)
	fmt.Println(&j)

	n()
	m()
	n()

	bbb = "G"
	print(bbb)
	f1()
}

func n() { print(aa) }

func m() {
	aa = "O"
	print(aa)
}

func f1() {
	bbb := "0"
	print(bbb)
	f2()
}

func f2() {
	print(bbb)
}
