package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

type Test struct {
	Name string
}

func main() {
	c := cron.New()
	_, _ = c.AddFunc("* */10 * * *", func() {
		fmt.Println(time.Now().Unix())
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
