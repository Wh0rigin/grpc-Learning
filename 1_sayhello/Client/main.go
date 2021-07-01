package main

import (
	"context"
	"fmt"
	"sayhello/services"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	c := services.NewBlackBoxClient(conn)

	req := services.Request{
		Id: 1,
	}

	reply, _ := c.SayHello(context.Background(), &req)

	fmt.Println(reply.Msg)
}
