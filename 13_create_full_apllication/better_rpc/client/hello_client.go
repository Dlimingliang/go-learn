package main

import (
	"fmt"
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/better_rpc/client_proxy"
)

func main() {
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply string
	client.Hello("lml", &reply)
	fmt.Println(reply)
}
