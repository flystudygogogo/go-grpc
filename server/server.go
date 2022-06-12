package main

import (
	"context"
	"fmt"

	"go-grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

//定义服务端 实现 约定的接口
type UserInfoService struct{}

var u = UserInfoService{}

//实现 interface
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *proto.UserRequest) (resp *proto.UserResponse, err error) {
	name := req.Name
	if name == "YMX" {
		resp = &proto.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Title: []string{"Java", "Go"},
		}
	}
	err = nil
	return
}
func main() {
	//1 添加监听的端口
	port := ":6666"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("端口监听错误 : %v\n", err)
	}
	fmt.Printf("正在监听： %s 端口\n", port)
	//2 启动grpc服务
	s := grpc.NewServer()
	//3 将UserInfoService服务注册到gRPC中
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量，需要取地址传参
	proto.RegisterUserInfoServiceServer(s, &u)
	s.Serve(l)
}
