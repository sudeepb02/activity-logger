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

	//Create a user menu to get input from user and pass the inputs to logUserActivity
	// logPlayTime(cli)
}

// func logUserActivity(string type, string timestamp, string duration, string label)

func logUserActivity(cli activitypb.ActivityServiceClient, activityType string, activityTimestamp string, duration string, label string) {

	fmt.Println("Logging activity for PlayTime")
	req := &activitypb.LogActivityRequest{
		Activity: &activitypb.Activity{
			Type: activityType,
			Timestamp: activityTimestamp,
			Duration: duration,
			Label: label,
		},
	}

	res, err := cli.LogActivity(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to log activity %v", err)
	}
	fmt.Printf("Response from server : %v", res.Result)
}