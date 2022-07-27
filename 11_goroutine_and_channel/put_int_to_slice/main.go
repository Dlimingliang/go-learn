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

func main() {
	start := time.Now()
	defer fmt.Println("总耗时:", time.Since(start))
	//addToSlice()
	addToSliceWaitGroup()
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
