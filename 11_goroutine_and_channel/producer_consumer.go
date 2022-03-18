package main

import (
	"fmt"
	"time"
)

type producer struct {
	data chan int
	quit chan chan error
}

func (p *producer) close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func main() {

	//ch := make(chan int, 10)
	//quit := make(chan interface{})
	//
	//go func() {
	//	var i = 0
	//	for {
	//		i = caculateNextInt(i)
	//		select {
	//		case ch <- i:
	//		case <-quit:
	//			close(ch)
	//			return
	//		}
	//
	//	}
	//}()
	//
	//for value := range ch{
	//	fmt.Printf("i=%v\n", value)
	//	if value >= 5 {
	//		close(quit)
	//	}
	//}

	prod := &producer{data: make(chan int, 10), quit: make(chan chan error)}

	go func() {
		var i = 0
		for {
			i = caculateNextInt(i)
			select {
			case prod.data <- i:
			case ch := <-prod.quit:
				close(prod.data)
				close(ch)
				return
			}
		}
	}()

	for i := range prod.data {
		fmt.Printf("consumer receive value=%v\n", i)
		if i >= 15 {
			err := prod.close()
			if err != nil {
				fmt.Printf("unexpected error: %v\n", err)
			}
		}
	}
}

func caculateNextInt(i int) int {
	time.Sleep(1 * time.Second)
	return i + 1
}
