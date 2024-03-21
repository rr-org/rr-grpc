package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"rr-grpc/winner"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type myWinnerServer struct {
	winner.UnimplementedWinnerServer
}

func (s myWinnerServer) Update(ctx context.Context, req *winner.CreateRequest) (*winner.CreateResponse, error) {

	if req.Id == "" {
		return nil, errors.New("Id is empty")
	}

	// load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// get env variables
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	// connect to db
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// disconnect
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// define db name & collection
	coll := client.Database("Resonance-Riddle").Collection("users")

	// convert to objectID
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	// find user
	var user User
	err = coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the id %s\n", req.Id)
		return nil, errors.New("User not found")
	}
	// add 1 diamond
	user.Diamond++

	// update to db
	update := bson.M{"$set": bson.M{"diamond": user.Diamond}}
	_, err = coll.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	return &winner.CreateResponse{
		Response: "success",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegister := grpc.NewServer()
	service := &myWinnerServer{}

	winner.RegisterWinnerServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Username  *string            `bson:"username" json:"username"`
	Avatar    *string            `bson:"avatar" json:"avatar"`
	Diamond   int                `bson:"diamond" json:"diamond"`
	Score     int                `bson:"score" json:"score"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
