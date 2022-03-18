package client_proxy

import (
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/better_rpc/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protcol, address string) HelloServiceStub {
	conn, _ := rpc.Dial(protcol, address)
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(param string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", param, reply)
	return err
}
