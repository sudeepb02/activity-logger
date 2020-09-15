package main

import (
	"log"
	"fmt"
	"context"
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

	logPlayTime(cli)
}

func logPlayTime(cli activitypb.ActivityServiceClient) {

	fmt.Println("Logging activity for PlayTime")
	req := &activitypb.PlayRequest{
		Activity: &activitypb.Activity{
			Type: "PlayTime",
			Timestamp: "Test",
			Duration: "1",
			Label: "TestPlay",
		},
	}

	res, err := cli.Play(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to log activity %v", err)
	}
	fmt.Printf("Response from server : %v", res.Result)
}