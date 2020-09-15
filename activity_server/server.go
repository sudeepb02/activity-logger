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

func (*server) Play(ctx context.Context, req *activitypb.PlayRequest) (*activitypb.PlayResponse, error) {
	fmt.Printf("Received request to log Play activity %v", req)
	activity := req.GetActivity()
	duration := activity.GetDuration()
	label := activity.GetLabel()

	res := &activitypb.PlayResponse{
		Result: "Logged activity " + label + " for duration of " + duration,
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