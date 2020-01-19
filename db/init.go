package db

import (
	"context"
	"time"

	"github.com/Splash07/nc_student/config"
	"github.com/Splash07/nc_student/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

const (
	DbName     = "GolangTest"
	ColName    = "Students"
	DbNameLog  = "Log"
	ColNameLog = "Log"
)

func init() {
	connect()
}

func connect() {
	clientOptions := options.Client().ApplyURI(config.Config.Mongo.Uri)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(4)
	clientOptions.SetReadPreference(readpref.Nearest())
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		utils.ErrorLogger.Fatalf("connect error :%v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		utils.ErrorLogger.Fatalf("ping error :%v", err)
	}
	Client = client
}
