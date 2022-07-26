package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//chanWait()
	waitGroupWait()
}

func waitGroupWait() {
	var wg sync.WaitGroup
	wg.Add(2)
	//go say(1, &wg)
	//go say(2, &wg)
	for i := 0; i < 3; i++ {
		go func(d int) {
			defer wg.Done()
			fmt.Printf("hello world%d\n", d)
		}(i)
	}
	wg.Wait()
	fmt.Println("main exited")
}

func say(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("hello world%d,\n", i)
}

func chanWait() {
	ch := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("hello world%d,\n", i)
		}
		close(ch)
	}()
	<-ch
	fmt.Println("main exited")
}
