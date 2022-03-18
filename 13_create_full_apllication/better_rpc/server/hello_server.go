package main

import (
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/better_rpc/handler"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (helloService *HelloService) Hello(param string, reply *string) error {
	*reply = "hello " + param
	return nil
}

func main() {
	//监听服务
	listener, _ := net.Listen("tcp", ":1234")
	//注册rpc
	rpc.RegisterName(handler.HelloServiceName, &HelloService{})
	//启动服务
	for {
		connection, _ := listener.Accept()
		go rpc.ServeConn(connection)
	}
}
