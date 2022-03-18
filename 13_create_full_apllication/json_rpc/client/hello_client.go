package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	connection, _ := net.Dial("tcp", "localhost:1234")
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(connection))
	var reply string
	client.Call("HelloService.Hello", "lml", &reply)
	fmt.Println(reply)
}
