package main

import (
	"fmt"
	"sync"
)

const producerCount int = 4
const consumerCount int = 3

var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

var multiMessages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

func singleProducer(data chan string) {
	for _, message := range messages {
		data <- message
	}
	close(data)
}

func multiProducer(data chan<- string, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, message := range multiMessages[idx] {
		fmt.Printf("Producer %v sending message \"%v\"\n", idx, message)
		data <- message
	}
}

func singleConsumer(data chan string, done chan bool) {
	for value := range data {
		fmt.Printf("Consumed message \"%v\"\n", value)
	}
	done <- true
}

//func multiConsumer(worker int, data chan string, done chan bool) {
//	for value := range data {
//		fmt.Printf("Message %v is consumed by worker %v.\n", value, worker)
//	}
//	done <- true
//}

func multiConsumer(data <-chan string, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range data {
		fmt.Printf("Message %v is consumed by worker %v.\n", value, id)
	}
}

func main() {

	//// single producer -> single consumer or multi consumer
	//data := make(chan string)
	//done := make(chan bool)
	//go singleProducer(data)
	//
	////go singleConsumer(data, done)
	//
	//for i := 1; i <= consumerCount; i++ {
	//	go multiConsumer(i, data, done)
	//}
	//<-done

	//// multi producer -> single consumer
	//data := make(chan string)
	//done := make(chan bool)
	//wg := sync.WaitGroup{}
	//
	//for i := 0; i < producerCount; i++ {
	//	wg.Add(1)
	//	go multiProducer(data, i, &wg)
	//}
	//
	//go singleConsumer(data, done)
	//
	//wg.Wait()
	//close(data)
	//<-done

	// multi producer -> multi consumer
	data := make(chan string)
	wp := &sync.WaitGroup{}
	wc := &sync.WaitGroup{}

	wp.Add(producerCount)
	wc.Add(consumerCount)

	for i := 0; i < producerCount; i++ {
		go multiProducer(data, i, wp)
	}

	for i := 0; i < consumerCount; i++ {
		go multiConsumer(data, i, wc)
	}

	wp.Wait()
	close(data)
	wc.Wait()
}
