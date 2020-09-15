package main

import (
	"log"
	"fmt"
	"google.golang.org/grpc"
	"github.com/sudeepb02/activity-logger/activitypb"
)

func main() {
	fmt.Println("Client started...")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to set Client connection %v ", err)
	}
	defer cc.Close()

	cli := activitypb.NewActivityServiceClient(cc)

	fmt.Printf("Client connection ready %v", cli)

}