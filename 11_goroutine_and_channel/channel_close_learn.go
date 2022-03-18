package main

import (
	"fmt"
	"time"
)

func main() {
	//通过关闭，是发送方，发送close. 接收方如何知道通道是否关闭,使用input, open := <-ch, open=true，通道是打开的。select来判断通道是否阻塞.
	ch := make(chan string)
	go sendString(ch)
	getString(ch)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)
}

func sendString(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch) //关闭通道
}

func getString(ch chan string) {
	//for {
	//	input, open := <-ch
	//	if !open {
	//		break
	//	}
	//	fmt.Println(input)
	//}
	for input := range ch {
		fmt.Println(input)
	}
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
