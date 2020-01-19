package db

import (
	"context"
	"time"

	"github.com/Splash07/nc_student/model"
	"github.com/Splash07/nc_student/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func AddStudent(student *model.Student) (interface{}, error) {
	var count []model.Count
	collection := Client.Database(DbName).Collection(ColName)
	collectionCount := Client.Database(DbName).Collection("count")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cur, err := collectionCount.Find(ctx, bson.M{})
	if err != nil {
		utils.ErrorLogger.Fatal(err)
	}
	cur.All(ctx, &count)
	student.ID = count[0].Value + 1
	filter := bson.M{"id": count[0].ID}
	update := bson.M{"$set": bson.M{"val": student.ID}}
	_, err = collectionCount.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.ErrorLogger.Printf("Update counter unsucessful %v", err)
	}

	res, err := collection.InsertOne(ctx, &student)
	return res, err
}

func UpdateStudent(student *model.StudentUpdateRequest) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": student.Email}
	update := bson.M{"$set": student}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}

func AddLog(log *model.Log) (interface{}, error) {
	collection := Client.Database(DbNameLog).Collection(ColNameLog)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, &log)
	if err != nil {
		utils.ErrorLogger.Printf("Insertion unsucessful %v", err)
	}
	return res, err
}
