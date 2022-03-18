package main

import "sync"

type Info struct {
	mu  sync.Mutex
	str string
}

func main() {

}

func update(info *Info, str string) {
	info.mu.Lock()
	info.str = str
	info.mu.Unlock()
}
