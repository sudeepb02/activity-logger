package main

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"fmt"
	"context"
	// "time"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"	

	"github.com/sudeepb02/activity-logger/activitypb"	

)

var collection *mongo.Collection

type server struct {

}

type activityItem struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Type string	`bson:"type"`
	Timestamp string `bson:"timestamp"`
	Duration string `bson:"duration"`
	Label string `bson:"label"`
}

func (*server) LogActivity(ctx context.Context, req *activitypb.LogActivityRequest) (*activitypb.LogActivityResponse, error) {
	fmt.Printf("Received request to log activity %v", req)
	activity := req.GetActivity()

	data := activityItem {
		Type: activity.GetType(),
		Timestamp: activity.GetTimestamp(),
		Duration: activity.GetDuration(),
		Label: activity.GetLabel(),
	}

	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error"),
		)

	}

	return &activitypb.LogActivityResponse{
		Result: "Activity logged successfully with ID " + oid.Hex(),
	}, nil
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

func (*server) AddUser(ctx context.Context, req *activitypb.AddUserRequest) (*activitypb.AddUserResponse, error) {
	fmt.Printf("Request received to add new user %v", req)
	// userDetails := req.GetUser()

	//Add user to DB
	return &activitypb.AddUserResponse{
		Result: true,
	}, nil
}

func main() {

	fmt.Println("Connecting to MongoDB")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB %v", err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalf("Error connecting to MongoDB %v", err)
	}

	fmt.Println("Activity Service Started")
	collection = client.Database("activitydb").Collection("activity")

/////////////////////////////////////////////////
//Mongo//

	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
////////////////////////////////////////////////////////	
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