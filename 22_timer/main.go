package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

type Test struct {
	Name string
}

func main() {
	c := cron.New()
	_ = c.AddFunc("*/1 * * * * ?", func() {
		var point *Test
		fmt.Println(point.Name)
	})
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
