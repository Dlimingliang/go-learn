package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1 * time.Second)

	chWait := make(chan string, 10)
	go longSendData(chWait)
	str := <-chWait
	fmt.Println(str)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func longSendData(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "longSendData"
}

func getData(ch chan string) {
	var input string
	time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Println(input)
	}
}
