package main

import (
	"fmt"
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/rpc/helloworld"
	"github.com/golang/protobuf/proto"
)

func main() {
	req := helloworld.HelloRequest{
		Name:    "lml",
		Age:     12,
		Courses: []string{"go", "gin"},
	}
	rsp, _ := proto.Marshal(&req)
	fmt.Println(string(rsp))
	newRsp := helloworld.HelloRequest{}
	proto.Unmarshal(rsp, &newRsp)
	fmt.Println(newRsp.Name, newRsp.Age, newRsp.Courses)
}
