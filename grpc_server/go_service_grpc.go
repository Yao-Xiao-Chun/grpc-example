package main

import (
	"context"
	"fmt"
	pb "go_work/grpc-go-tools/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// server 定一个空的结构体
type server struct {
	pb.UnimplementedGreeterServer
}



func (s *server) SayHello(ctx context.Context,in *pb.ImageReq) (out *pb.ImageRequestResp,err error) {

	return &pb.ImageRequestResp{Message: "hello GRPC"},nil
}

// main grpc service
func main()  {

	fmt.Println("GRPC:start...")
	// 监听127.0.0.1:50051地址
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	pb.RegisterGreeterServer(s, &server{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	fmt.Println("GRPC:127.0.0.1:50051")

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}


