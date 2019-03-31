package main

import (
	"log"
	"net"

	pb "github.com/niemuuu/gRPC-sample/pb"
	service "github.com/niemuuu/gRPC-sample/service"

	"google.golang.org/grpc"
)

func main() {
	// 19003番ポートでlisten
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &service.MyCatService{}

	// CatServiceの登録
	pb.RegisterCatServer(server, catService)

	// serve
	server.Serve(listenPort)
}
