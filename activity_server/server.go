package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gopkg.in/mgo.v2/bson"

	"github.com/sudeepb02/activity-logger/activitypb"
)

var collection *mongo.Collection

type server struct {
}

type activityItem struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    string             `bson:"user_id"`
	Type      string             `bson:"type"`
	Timestamp int64              `bson:"timestamp"`
	Duration  int64              `bson:"duration"`
	Label     string             `bson:"label"`
}

type userItem struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone string             `bson:"phone"`
}

func (*server) LogActivity(ctx context.Context, req *activitypb.LogActivityRequest) (*activitypb.LogActivityResponse, error) {

	fmt.Printf("Received request to log user activity %v \n", req)

	activity := req.GetActivity()
	data := activityItem{
		UserID:    activity.GetUserId(),
		Type:      activity.GetType(),
		Timestamp: activity.GetTimestamp(),
		Duration:  activity.GetDuration(),
		Label:     activity.GetLabel(),
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
		Activity: &activitypb.Activity{
			Id:        oid.Hex(),
			UserId:    activity.GetUserId(),
			Type:      activity.GetType(),
			Timestamp: activity.GetTimestamp(),
			Duration:  activity.GetDuration(),
			Label:     activity.GetLabel(),
		},
	}, nil
}

func (*server) IsDone(ctx context.Context, req *activitypb.IsDoneRequest) (*activitypb.IsDoneResponse, error) {

	fmt.Printf("Request received at server to check if activity is completed %v", req)

	getActivityReq := &activitypb.GetActivityRequest{
		Id: req.GetId(),
	}

	activityRes, err := (*server).GetActivity(&server{}, context.Background(), getActivityReq)
	if err != nil {
		log.Fatalf("Error getting activity details %v", err)
	}

	activityDetails := activityRes.Activity
	currentTimestamp := time.Now().Unix()
	status := false

	if activityDetails.Timestamp+activityDetails.Duration < currentTimestamp {
		status = true
	}

	res := &activitypb.IsDoneResponse{
		Status: status,
	}
	return res, nil
}

func (*server) IsValid(ctx context.Context, req *activitypb.IsValidRequest) (*activitypb.IsValidResponse, error) {

	fmt.Printf("Request received at server to check if activity is valid %v", req)

	getActivityReq := &activitypb.GetActivityRequest{
		Id: req.GetId(),
	}

	activityRes, err := (*server).GetActivity(&server{}, context.Background(), getActivityReq)
	if err != nil {
		log.Fatalf("Error getting activity details %v", err)
	}

	userDetailsReq := &activitypb.GetUserRequest{
		Id: activityRes.Activity.UserId,
	}

	_, getUserErr := (*server).GetUser(&server{}, context.Background(), userDetailsReq)
	if err != nil {
		return &activitypb.IsValidResponse{
			Result: false,
		}, getUserErr
	}

	return &activitypb.IsValidResponse{
		Result: true,
	}, nil
}

func (*server) AddUser(ctx context.Context, req *activitypb.AddUserRequest) (*activitypb.AddUserResponse, error) {

	fmt.Printf("Request received to add new user %v \n", req)
	userDetails := req.GetUser()

	//Add user to DB
	data := userItem{
		Name:  userDetails.GetName(),
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
		User: &activitypb.User{
			Id:    oid.Hex(),
			Name:  userDetails.GetName(),
			Email: userDetails.GetEmail(),
			Phone: userDetails.GetPhone(),
		},
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
	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find user %v", err),
		)
	}

	return &activitypb.GetUserResponse{
		User: &activitypb.User{
			Id:    data.ID.Hex(),
			Name:  data.Name,
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
	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find activity %v", err),
		)
	}

	return &activitypb.GetActivityResponse{
		Activity: &activitypb.Activity{
			Id:        data.ID.Hex(),
			Type:      data.Type,
			Timestamp: data.Timestamp,
			Duration:  data.Duration,
			Label:     data.Label,
		},
	}, nil
}

func (*server) UpdateActivity(ctx context.Context, req *activitypb.UpdateActivityRequest) (*activitypb.UpdateActivityResponse, error) {

	fmt.Printf("Request received to update activity %v \n", req)
	activity := req.GetActivity()
	oid, err := primitive.ObjectIDFromHex(activity.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot parse ID"),
		)
	}
	data := &activityItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(context.Background(), filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find activity %v", err),
		)
	}
	data.UserID = activity.GetUserId()
	data.Type = activity.GetType()
	data.Timestamp = activity.GetTimestamp()
	data.Duration = activity.GetDuration()
	data.Label = activity.GetLabel()

	_, updErr := collection.ReplaceOne(context.Background(), filter, data)
	if updErr != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error updating activity %v", updErr),
		)
	}

	return &activitypb.UpdateActivityResponse{
		Activity: &activitypb.Activity{
			Id:        data.ID.Hex(),
			UserId:    data.UserID,
			Type:      data.Type,
			Timestamp: data.Timestamp,
			Duration:  data.Duration,
			Label:     data.Label,
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
