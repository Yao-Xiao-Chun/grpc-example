package main

import (
	"context"
	"fmt"
	pb "go_work/grpc-go-tools/proto"
	"google.golang.org/grpc"
	"time"
)

// main  grpc调用
func main()  {

	var localAddr = "127.0.0.1" //ip:port
	var port = ":50051"

	fmt.Println("conn grpc start....")

	conn,err := grpc.Dial(localAddr+port,grpc.WithInsecure())

	if err != nil{
		fmt.Println("连接错误",err.Error())
		return
	}
	defer conn.Close()

	//初始化服务端
	c := pb.NewGreeterClient(conn)

	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	r,err := c.SayHello(ctx,&pb.ImageReq{
		Name: "",
	})

	if err != nil{
		fmt.Println("调用错误")
	}

	fmt.Println(r.Message)

}