package main

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"github.com/sudeepb02/activity-logger/activitypb"
)

type server struct {

}

func main() {

	fmt.Println("Server started, listening on Port 50051...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Error listening on Port 50051 %v", err)
	}

	ser := grpc.NewServer()

	activitypb.RegisterActivityServiceServer(ser, &server{})

	if err := ser.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on Port 50051 %v", err)
	}
}