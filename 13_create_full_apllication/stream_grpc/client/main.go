package main

import (
	"context"
	"fmt"
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/stream_grpc/stream"
	"google.golang.org/grpc"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	call := stream.NewGreetereClient(conn)
	//res, _ := call.GetStream(context.Background(), &stream.StreamReq{
	//	Data: "lml",
	//})
	//for {
	//	value, err := res.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(value)
	//}
	//
	//i := 0
	//putS, _ := call.PutStream(context.Background())
	//for {
	//	i++
	//	_ = putS.Send(&stream.StreamReq{
	//		Data: fmt.Sprintf("lml%d", i),
	//	})
	//	if i > 10 {
	//		break
	//	}
	//}

	ws := sync.WaitGroup{}
	ws.Add(2)
	allStr, _ := call.AllStream(context.Background())
	go func() {
		defer ws.Done()
		for true {
			if value, err := allStr.Recv(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("收到服务端消息", value.Data)
			}
		}
	}()

	go func() {
		defer ws.Done()
		for true {
			_ = allStr.Send(&stream.StreamReq{
				Data: "我是客户端",
			})
			time.Sleep(time.Second)
		}
	}()
	ws.Wait()
}
