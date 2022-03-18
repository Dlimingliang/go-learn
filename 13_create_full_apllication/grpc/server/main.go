package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Dlimingliang/go-learn/13_create_full_apllication/grpc/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

type Server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if nameSlice, ok := md["name"]; ok {
			for k, v := range nameSlice {
				fmt.Println(k, v)
			}
		}
	}

	return &helloworld.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

func main() {

	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		panic(err.Error())
	}

	//注册服务端耗时Interceptor
	//timeInterceptorFunc := func(ctx context.Context, req interface{},
	//	info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	//	now := time2.Now()
	//	res, err := handler(ctx, req)
	//	fmt.Printf("服务端耗时%s\n", time2.Since(now))
	//	return res, err
	//}
	//timeInterceptor := grpc.UnaryInterceptor(timeInterceptorFunc)

	//注册服务端认证Interceptor
	authInterceptorFunc := func(ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		//time.Sleep(time.Second * 5)

		md, ok := metadata.FromIncomingContext(ctx)
		var appid string
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token信息")
		}
		if nameSlice, ok := md["appid"]; ok {
			appid = nameSlice[0]
		}
		if appid != "111" {
			return resp, status.Error(codes.Unauthenticated, "token不正确")
		}

		res, err := handler(ctx, req)
		return res, err
	}
	authInterceptor := grpc.UnaryInterceptor(authInterceptorFunc)

	var serverOptions []grpc.ServerOption
	//serverOptions = append(serverOptions, timeInterceptor)
	serverOptions = append(serverOptions, authInterceptor)
	s := grpc.NewServer(serverOptions...)
	helloworld.RegisterGreeterServer(s, &Server{})

	err = s.Serve(listen)
	if err != nil {
		panic(err.Error())
	}
}
