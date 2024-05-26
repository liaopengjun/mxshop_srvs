package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/initialize"
	"mxshop_srvs/user_srv/proto"
	"net"
)

func main() {
	// 初始化
	initialize.Viper()
	initialize.InitDB()
	initialize.InitLogger()
	zap.S().Info(global.ServerConfig)

	// 启动grpc
	server := grpc.NewServer()
	// 注册服务
	proto.RegisterUserServer(server, &handler.UserServer{})
	//监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	fmt.Println("grpc server listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}

}
