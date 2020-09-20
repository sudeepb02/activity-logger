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

	//User Menu
	fmt.Println("Please select an option: ")
	fmt.Println("1. Log user activity\n2. Add user\n3. Get User\n4. Get Activity\n5. Update activity\n6. Check activity status\n7. Check activity validity")

	var i int
	fmt.Scanf("%d", &i)
	fmt.Printf("Option selected by user: %d\n", i)	

	switch i {
	case 1 :
		logUserActivity(cli)
	case 2 : 
		addUser(cli)
	case 3 :
		getUser(cli)
	case 4 : 
		getActivity(cli)
	case 5 :
		updateActivity(cli)
	case 6 :
		checkActivityStatus(cli)
	case 7 :
		checkActivityValidity(cli)
	default :
		fmt.Println("Please select a correct option")		
	}	
}

func updateActivity(cli activitypb.ActivityServiceClient) {

	var id, activityType, activityLabel string
	var activityTimestamp, activityDuration int64

	fmt.Println("Enter activity ID:")
	fmt.Scanln(&id)

	fmt.Println("Enter activity type:")
	fmt.Scanln(&activityType)

	fmt.Println("Enter the timestamp:")
	fmt.Scanln(&activityTimestamp)

	fmt.Println("Enter duration of the activity")
	fmt.Scanln(&activityDuration)

	fmt.Println("Enter label for the activity:")
	fmt.Scanln(&activityLabel)

	fmt.Println("Updating User activity...")
	req := &activitypb.UpdateActivityRequest{
		Activity: &activitypb.Activity{
			Id: id,
			Type: activityType,
			Timestamp: activityTimestamp,
			Duration: activityDuration,
			Label: activityLabel,
		},
	}

	res, err := cli.UpdateActivity(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to update activity %v", err)
	}
	fmt.Printf("Activity updated successfully: %v \n", res.Activity)	

}

func getUser(cli activitypb.ActivityServiceClient) {
	
	var userID string
	fmt.Print("Enter user ID to get user details: ")
	fmt.Scanln(&userID)

	req := &activitypb.GetUserRequest {
		Id: userID,
	}

	fmt.Println("Getting user details...")
	res, err := cli.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get user details %v", err)
	}
	fmt.Printf("Response : %v \n", res.User)
}

func getActivity(cli activitypb.ActivityServiceClient) {
	
	var activityID string
	fmt.Print("Enter activity ID to get details: ")
	fmt.Scanln(&activityID)

	req := &activitypb.GetActivityRequest {
		Id: activityID,
	}

	fmt.Println("Getting activity details...")
	res, err := cli.GetActivity(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get user details %v", err)
	}
	fmt.Printf("Response : %v \n", res.Activity)
}

func addUser(cli activitypb.ActivityServiceClient) {

	var name, email, phone string

	fmt.Print("Enter name of the User: ")
	fmt.Scanln(&name)

	fmt.Print("Enter email address: ")
	fmt.Scanln(&email)

	fmt.Print("Enter Phone number: ")
	fmt.Scanln(&phone)

	req := &activitypb.AddUserRequest {
		User: &activitypb.User {
			Name: name,
			Email: email,
			Phone: phone,
		},
	}

	fmt.Println("Adding a new user...")	
	res, err := cli.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to add user %v", err)
	}
	fmt.Printf("Response : %v \n", res.Result)
}

func logUserActivity(cli activitypb.ActivityServiceClient) {

	var activityType, activityLabel string
	var activityTimestamp, activityDuration int64

	fmt.Println("Enter activity type:")
	fmt.Scanln(&activityType)

	fmt.Println("Enter the timestamp:")
	fmt.Scanln(&activityTimestamp)

	fmt.Println("Enter duration of the activity")
	fmt.Scanln(&activityDuration)

	fmt.Println("Enter label for the activity:")
	fmt.Scanln(&activityLabel)

	fmt.Println("Logging User activity...")
	req := &activitypb.LogActivityRequest{
		Activity: &activitypb.Activity{
			Type: activityType,
			Timestamp: activityTimestamp,
			Duration: activityDuration,
			Label: activityLabel,
		},
	}

	res, err := cli.LogActivity(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to log activity %v", err)
	}
	fmt.Printf("Response from server : %v \n", res.Result)
}

func checkActivityStatus(cli activitypb.ActivityServiceClient) {
	
	// fmt.Println("Checking activity status...")
	var activityID string

	fmt.Println("Please enter activity ID: ")
	fmt.Scanln(&activityID)

	req := &activitypb.IsDoneRequest{
		Id : activityID,
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

func checkActivityValidity(cli activitypb.ActivityServiceClient) {
	
	var activityID string

	fmt.Println("Please enter activity ID: ")
	fmt.Scanln(&activityID)

	req := &activitypb.IsValidRequest{
		Id : activityID,
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