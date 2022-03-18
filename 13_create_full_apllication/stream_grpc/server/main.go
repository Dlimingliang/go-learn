package main

import (
	"fmt"
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/stream_grpc/stream"
	"google.golang.org/grpc"
	"net"
	"sync"
	"time"
)

type Server struct {
	stream.UnimplementedGreetereServer
}

func (s *Server) GetStream(req *stream.StreamReq, resStr stream.Greetere_GetStreamServer) error {

	i := 0
	for {
		i++
		_ = resStr.Send(&stream.StreamRes{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}
func (s *Server) PutStream(cliStr stream.Greetere_PutStreamServer) error {

	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

func (s *Server) AllStream(allStr stream.Greetere_AllStreamServer) error {

	ws := sync.WaitGroup{}
	ws.Add(2)
	go func() {
		defer ws.Done()
		for true {
			if value, err := allStr.Recv(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("收到客户端消息", value.Data)
			}
		}
	}()

	go func() {
		defer ws.Done()
		for true {
			_ = allStr.Send(&stream.StreamRes{
				Data: "我是服务端",
			})
			time.Sleep(time.Second)
		}
	}()
	ws.Wait()
	return nil
}

func main() {

	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	stream.RegisterGreetereServer(s, &Server{})

	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
