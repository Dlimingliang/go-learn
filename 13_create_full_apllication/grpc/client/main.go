package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/grpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	time2 "time"
)

const (
	defaultName = "World"
)

var (
	addr = flag.String("addr", "localhost:9090", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type customCredentials struct {
}

func (c customCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid": "11",
	}, nil
}
func (c customCredentials) RequireTransportSecurity() bool {
	return false
}

func main() {

	flag.Parse()

	//设置客户端耗时拦截器
	timeInterceptorFunc := func(ctx context.Context, method string, req, reply interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		now := time2.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("客户端调用耗时%s\n", time2.Since(now))
		return err
	}
	timeInterceptor := grpc.WithUnaryInterceptor(timeInterceptorFunc)

	//设置认证拦截器
	authInterceptor := grpc.WithPerRPCCredentials(customCredentials{})

	var dialOptions []grpc.DialOption
	dialOptions = append(dialOptions, timeInterceptor)
	dialOptions = append(dialOptions, authInterceptor)
	dialOptions = append(dialOptions, grpc.WithInsecure())
	conn, err := grpc.Dial(*addr, dialOptions...)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	//设置元数据
	md := metadata.New(map[string]string{
		"name":     "lml",
		"password": "jiale",
	})
	//设置超时机制
	context, _ := context.WithTimeout(context.Background(), time2.Second*3)
	ctx := metadata.NewOutgoingContext(context, md)
	call := helloworld.NewGreeterClient(conn)
	result, err := call.SayHello(ctx, &helloworld.HelloRequest{
		Name: *name,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			panic(err)
		}
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
	if result != nil {
		fmt.Println("Greeting: ", result.Message)
	}
}
