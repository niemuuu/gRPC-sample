package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/niemuuu/gRPC-sample/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("client connection error", err)
	}
	defer conn.Close()
	client := pb.NewCatClient(conn)
	message := &pb.GetMyCatMessage{TargetCat: "mike"}
	res, err := client.GetMyCat(context.TODO(), message)
	fmt.Println("result:%#v \n", res)
	fmt.Println("error::%#v \n", err)
}
