package main

import (
	"log"
	"net"
	"fmt"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"	

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"	

	"gopkg.in/mgo.v2/bson"

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

type userItem struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	Phone string `bson:"phone"`
}

func (*server) LogActivity(ctx context.Context, req *activitypb.LogActivityRequest) (*activitypb.LogActivityResponse, error) {
	fmt.Printf("Received request to log activity %v \n", req)
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
	fmt.Printf("Request received to add new user %v \n", req)
	userDetails := req.GetUser()

	//Add user to DB
	data := userItem{
		Name: userDetails.GetName(),
		Email: userDetails.GetEmail(),
		Phone: userDetails.GetPhone(),
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

	return &activitypb.AddUserResponse{
		Result: "User added successfully with ID " + oid.Hex(),
	}, nil
}

func (*server) GetUser(ctx context.Context, req *activitypb.GetUserRequest) (*activitypb.GetUserResponse, error) {
	fmt.Printf("Request received to get user details %v \n", req)
	userID := req.GetId()
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v", err),
		)
	}

	data := &userItem{}
	filter := bson.M{ "_id": oid }

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user %v", err),
		)
	}

	return &activitypb.GetUserResponse{
		User: &activitypb.User {
			Id: data.ID.Hex(),
			Name: data.Name,
			Email: data.Email,
			Phone: data.Phone,
		},
	}, nil
}

func (*server) GetActivity(ctx context.Context, req *activitypb.GetActivityRequest) (*activitypb.GetActivityResponse, error) {

	fmt.Printf("Request received to get activity %v \n", req)

	activityID := req.GetId()
	oid, err := primitive.ObjectIDFromHex(activityID)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v", err),
		)
	}

	data := &activityItem{}
	filter := bson.M{ "_id": oid }

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find activity %v", err),
		)
	}

	return &activitypb.GetActivityResponse{
		Activity: &activitypb.Activity {
			Id: data.ID.Hex(),
			Type: data.Type,
			Timestamp: data.Timestamp,
			Duration: data.Duration,
			Label: data.Label,
		},
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
	collection = client.Database("db").Collection("activity")
	
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