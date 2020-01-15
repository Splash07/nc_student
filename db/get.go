package db

import (
	"context"
	"encoding/json"
	mw "github.com/Splash07/nc_student/middleware"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllStudent() (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{} //map[string]interface{}
	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		mw.ErrorLogger.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		mw.ErrorLogger.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

func SearchStudentSimple(req StudentSearchRequest) (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var filter bson.M
	bs, err := json.Marshal(req)
	if err != nil {
		mw.ErrorLogger.Printf("Marshal error: %v", err)
	}

	err = json.Unmarshal(bs, &filter)
	if err != nil {
		mw.ErrorLogger.Printf("Unmarshal error: %v", err)
	}

	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		mw.ErrorLogger.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		mw.ErrorLogger.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

func SearchStudentCustom(req StudentSearchRequest) (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var query bson.D

	if req.FirstName != "" {
		query = append(query, bson.E{Key: "first_name", Value: bson.D{
			{"$regex", primitive.Regex{Pattern: req.FirstName, Options: "i"}},
		}})
	}

	if req.LastName != "" {
		query = append(query, bson.E{Key: "last_name", Value: bson.D{
			{"$regex", primitive.Regex{Pattern: req.LastName, Options: "i"}},
		}})
	}

	if req.Age != 0 {
		query = append(query, bson.E{Key: "age", Value: req.Age})
	}

	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, query)
	if err != nil {
		mw.ErrorLogger.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		mw.ErrorLogger.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}
