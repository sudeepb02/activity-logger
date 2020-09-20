package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	var userID string

	//User Menu
	fmt.Println("Please select an option: ")
	fmt.Println("1. Login as Existing user")
	fmt.Println("2. Register as a new user")
	fmt.Println("3. Log an activity")
	fmt.Println("4. Get User details")
	fmt.Println("5. Get Activity details")
	fmt.Println("6. Update activity")
	fmt.Println("7. Check activity status")
	fmt.Println("8. Check activity validity")
	fmt.Println("9. Exit")

	var i int
	fmt.Scanf("%d", &i)
	fmt.Printf("Option selected by user: %d\n", i)

	switch i {
	case 1:
		userLogin(cli, &userID)
	case 2:
		addUser(cli, &userID)
	case 3:
		logUserActivity(cli, &userID)
	case 4:
		getUser(cli)
	case 5:
		getActivity(cli)
	case 6:
		updateActivity(cli)
	case 7:
		checkActivityStatus(cli)
	case 8:
		checkActivityValidity(cli)
	case 9:
		break
	default:
		fmt.Println("Please select a correct option")
	}
}

func userLogin(cli activitypb.ActivityServiceClient, userID *string) {

	var id string
	fmt.Println("Please enter the user ID: ")
	fmt.Scanln(&id)

	req := &activitypb.GetUserRequest{
		Id: id,
	}

	fmt.Println("Logging in, please wait...")
	res, err := cli.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Login failed! Failed to get user details %v", err)
	}
	fmt.Printf("User logged in successfully : %v \n", res.User)
	userID = &id
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
			Id:        id,
			Type:      activityType,
			Timestamp: activityTimestamp,
			Duration:  activityDuration,
			Label:     activityLabel,
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

	req := &activitypb.GetUserRequest{
		Id: userID,
	}

	fmt.Println("Getting user details...")
	res, err := cli.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get user details %v", err)
	}
	fmt.Printf("User Details: \n%v \n", res.User)
}

func getActivity(cli activitypb.ActivityServiceClient) {

	var activityID string
	fmt.Print("Enter activity ID to get details: ")
	fmt.Scanln(&activityID)

	req := &activitypb.GetActivityRequest{
		Id: activityID,
	}

	fmt.Println("Getting activity details...")
	res, err := cli.GetActivity(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get user details %v", err)
	}
	fmt.Printf("Response : %v \n", res.Activity)
}

func addUser(cli activitypb.ActivityServiceClient, userID *string) {

	var name, email, phone string

	fmt.Print("Enter name of the User: ")
	fmt.Scanln(&name)

	fmt.Print("Enter email address: ")
	fmt.Scanln(&email)

	fmt.Print("Enter Phone number: ")
	fmt.Scanln(&phone)

	req := &activitypb.AddUserRequest{
		User: &activitypb.User{
			Name:  name,
			Email: email,
			Phone: phone,
		},
	}
	fmt.Println("Adding a new user...")

	res, err := cli.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to add user %v", err)
	}
	fmt.Printf("User added successfully. Logged in as registered user: %v \n", res.User)
	userID = &res.User.Id
}

func logUserActivity(cli activitypb.ActivityServiceClient, userID *string) {

	var activityType, activityLabel string
	var activityDuration int64

	fmt.Println("Enter activity type:")
	fmt.Scanln(&activityType)

	fmt.Println("Enter duration of the activity in seconds")
	fmt.Scanln(&activityDuration)

	fmt.Println("Enter label for the activity:")
	fmt.Scanln(&activityLabel)

	activityTimestamp := time.Now().Unix()

	fmt.Println("Logging User activity...")
	req := &activitypb.LogActivityRequest{
		Activity: &activitypb.Activity{
			UserId:    *userID,
			Type:      activityType,
			Timestamp: activityTimestamp,
			Duration:  activityDuration,
			Label:     activityLabel,
		},
	}

	res, err := cli.LogActivity(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to log activity %v", err)
	}
	fmt.Printf("Activity logged successfully : %v \n", res.Activity)
}

func checkActivityStatus(cli activitypb.ActivityServiceClient) {

	// fmt.Println("Checking activity status...")
	var activityID string

	fmt.Println("Please enter activity ID: ")
	fmt.Scanln(&activityID)

	req := &activitypb.IsDoneRequest{
		Id: activityID,
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
		Id: activityID,
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
