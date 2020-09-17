package main

import (
	"log"
	"net"
	"fmt"
	"context"
	"google.golang.org/grpc"
	"github.com/sudeepb02/activity-logger/activitypb"
)

type server struct {

}

func (*server) LogActivity(ctx context.Context, req *activitypb.LogActivityRequest) (*activitypb.LogActivityResponse, error) {
	fmt.Printf("Received request to log activity %v", req)
	activity := req.GetActivity()
	duration := activity.GetDuration()
	label := activity.GetLabel()

	res := &activitypb.LogActivityResponse{
		Result: "Logged activity " + label + " for duration of " + duration,
	}
	return res, nil
}

func (*server) IsDone(ctx context.Context, req *activitypb.IsDoneRequest) (*activitypb.IsDoneResponse, error) {
	fmt.Printf("Request received at server to check if activity is completed %v", req)
	// activityDetails := req.GetActivity()
	//Add logic to check if activity is complete, needs to convert types

	res := &activitypb.IsDoneResponse {
		Status: true,
	}

	return res, nil
}
	

func (*server) IsValid(ctx context.Context, req *activitypb.IsValidRequest) (*activitypb.IsValidResponse, error) {
	
	fmt.Printf("Request received at server to check if activity is completed %v", req)
	// activityDetails := req.GetActivity()
	//Add logic to check if activity is complete, needs to convert types

	res := &activitypb.IsValidResponse {
		Result: true,
	}

	return res, nil
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