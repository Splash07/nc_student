package db

import (
	"context"
	"time"

	"github.com/Splash07/nc_student/model"
	"github.com/Splash07/nc_student/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteStudentById(student *model.StudentDeleteRequest) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{"id": student.ID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		utils.ErrorLogger.Printf("Delete operation unsucessful %v", err)
	}
	return res, err
}
