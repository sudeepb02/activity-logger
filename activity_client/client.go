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
	var activityType string = "Play"
	var activityTimestamp string = "12345"
	var activityDuration string = "15"
	var activityLabel string = "PlayTime"

	logUserActivity(cli, activityType, activityTimestamp, activityDuration, activityLabel)

	//Create User
	name := "Sudeep"
	email := "sudeepbiswas02@gmail.com"
	phone := "123456789"

	addUser(cli, name, email, phone)
}

func addUser(cli activitypb.ActivityServiceClient, name string, email string, phone string) {
	fmt.Println("Adding a new user...")

	req := &activitypb.AddUserRequest {
		User: &activitypb.User {
			Name: name,
			Email: email,
			Phone: phone,
		},
	}

	res, err := cli.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to add user %v", err)
	}
	fmt.Printf("Response : %v \n", res)
}

func logUserActivity(cli activitypb.ActivityServiceClient, activityType string, activityTimestamp string, duration string, label string) {

	fmt.Println("Logging User activity...")
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
	fmt.Printf("Response from server : %v \n", res.Result)
}

func checkActivityStatus(cli activitypb.ActivityServiceClient, activityDetails *activitypb.Activity) {
	
	fmt.Println("Checking activity status...")
	req := &activitypb.IsDoneRequest{
		Activity : activityDetails,
	}

	res, err := cli.IsDone(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to check activity status %v", err)
	}
	if res.Status {
		fmt.Println("Activity is completed")
	} else {
		fmt.Println("Activity yet to be completed")
	}
}

func checkActivityValidity(cli activitypb.ActivityServiceClient, activityDetails *activitypb.Activity) {
	
	fmt.Println("Checking if the activity is valid...")
	req := &activitypb.IsValidRequest{
		Activity : activityDetails,
	}

	res, err := cli.IsValid(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to check activity %v", err)
	}
	if res.Result {
		fmt.Println("Activity is valid")
	} else {
		fmt.Println("Activity is invalid")
	}
}