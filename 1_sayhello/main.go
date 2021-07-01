package main

import (
	"net"
	"sayhello/services"

	"google.golang.org/grpc"
)

func main() {
	listener, _ := net.Listen("tcp", ":1234")
	myService := services.BlackBoxServices{}
	grpcServer := grpc.NewServer()
	services.RegisterBlackBoxServer(grpcServer, &myService)
	_ = grpcServer.Serve(listener)
}
