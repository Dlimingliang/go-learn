package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	slc  []int
	n    = 10000
	wg   sync.WaitGroup
	lock sync.Mutex
)

type Data struct {
	data []int
	ch   chan int
}

func (s *Data) Consumer() {
	for i := range s.ch {
		s.data = append(s.data, i)
	}
}

func (s *Data) Producer(value int) {
	s.ch <- value
}

func (s *Data) CloseChannel() {
	close(s.ch)
}

func NewScheduleJob(size int, done func()) *Data {
	s := &Data{
		ch:   make(chan int, size),
		data: make([]int, 0),
	}

	go func() {
		s.Consumer()
		done()
	}()
	return s
}

func addToSliceByChannel() {
	var (
		n  = 10000
		wg sync.WaitGroup
	)
	c := make(chan struct{})
	s := NewScheduleJob(n, func() { c <- struct{}{} })
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(value int) {
			defer wg.Done()
			s.Producer(value)
		}(i)
	}
	wg.Wait()
	s.CloseChannel()
	<-c
	fmt.Println(len(s.data))
}

func main() {
	start := time.Now()
	defer fmt.Println("总耗时:", time.Since(start))
	//addToSlice()
	//addToSliceWaitGroup()
	addToSliceByChannel()
}

func addToSlice() {
	for i := 0; i < n; i++ {
		slc = append(slc, i)
	}
	fmt.Println("len:", len(slc))
}

func addToSliceWaitGroup() {
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			slc = append(slc, i)
		}()
	}
	wg.Wait()
	fmt.Println("len:", len(slc))
}
