package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(1 * time.Second)
	boom := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
		default:
			fmt.Println("    .")
			time.Sleep(2 * time.Second)
		}
	}
}
