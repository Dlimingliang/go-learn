package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%02d", t.Day(), t.Month(), t.Year())
	fmt.Println()
	t = t.Add(60 * 60 * 24 * 1 * 1e9)
	fmt.Println(t.Day())
	time.Sleep(1 * time.Second)
}
