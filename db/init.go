package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Splash07/nc_student/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

var Client *mongo.Client

func Test() interface{} {
	fmt.Println("connect & insert db")
	return insertNunber()
}

func init() {

	connect()
}

func insertNunber() interface{} {
	collection := Client.Database("GolangTest").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		fmt.Println("Error occured!, ", err)
	}
	id := res.InsertedID
	return id
}

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.Uri))
	if err != nil {
		log.Fatalf("connect error :%v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Fatalf("connect error :%v", err)
	}
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	Client = client
}
