package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func AddStudent(student *Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, student)
	return res, err
}

func AddStudentAutoIncrement() {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"counterName": "counterStudents"}

	change := bson.M{"$inc": bson.M{"counterValue": 1}}

	err := collection.FindOneAndUpdate(ctx, filter, change)

	if err != nil {
		fmt.Println("Error occured when updating the counter !")
	}

	fmt.Println("Counter updated!!")

	//res, resErr := collection.InsertOne(ctx, student)

	//return res, resErr
}

func UpdateStudent(student *StudentUpdateRequest) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": student.Email}
	update := bson.M{"$set": student}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
