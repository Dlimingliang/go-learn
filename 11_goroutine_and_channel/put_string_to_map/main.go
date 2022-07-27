package main

import (
	"fmt"
	"sync"
	"time"
)

var strMap = map[int]byte{}
var lockMap *LockMap
var syncMap sync.Map

type LockMap struct {
	sync.RWMutex
	Map map[int]byte
}

func (l *LockMap) WriteMap(key int, value byte) {
	l.Lock()
	l.Map[key] = value
	l.Unlock()
}

func main() {
	start := time.Now()
	defer fmt.Println("总耗时:", time.Since(start).Nanoseconds())
	str := "xjioaoijfnqlerlulijasdxjioaoijfnqlerlulijasdxjioaoijfnqlerlulijasdxjioaoijfnqlerlulijasdxjioaoijfnqlerlulijasdxjioaoijfnqlerlulijasdxjioaoijfnqlerlulijasdxjioaoijfnqlerlulijasdf"
	//addToMap(str)
	//addToMapWaitGroup(str)
	addToSyncMapWaitGroup(str)
}

func addToMap(str string) {
	for i := 0; i < len(str); i++ {
		time.Sleep(100 * time.Millisecond)
		strMap[i] = str[i]
	}
	fmt.Println(strMap)
}

func addToMapWaitGroup(str string) {
	lockMap = &LockMap{
		Map: make(map[int]byte),
	}
	var mg sync.WaitGroup
	mg.Add(len(str))
	for i := 0; i < len(str); i++ {
		go func(index int, value byte) {
			defer mg.Done()
			time.Sleep(100 * time.Millisecond)
			lockMap.WriteMap(index, value)
		}(i, str[i])
	}
	mg.Wait()
	fmt.Println(lockMap.Map)
}

func addToSyncMapWaitGroup(str string) {
	var mg sync.WaitGroup
	mg.Add(len(str))
	for i := 0; i < len(str); i++ {
		go func(index int, value byte) {
			defer mg.Done()
			time.Sleep(100 * time.Millisecond)
			syncMap.Store(index, value)
		}(i, str[i])
	}
	mg.Wait()
	fmt.Println(syncMap)
}
