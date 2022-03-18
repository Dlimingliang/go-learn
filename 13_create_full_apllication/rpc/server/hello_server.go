package main

import (
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
	rpc.RegisterName("HelloService", &HelloService{})
	//启动服务
	connection, _ := listener.Accept()
	rpc.ServeConn(connection)
}
