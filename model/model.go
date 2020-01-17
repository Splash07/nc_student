package model

import "github.com/dgrijalva/jwt-go"

type StudentUpdateRequest struct {
	ID        int    `json:"id,omitempty" bson:"id"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	ClassName string `json:"class_name,omitempty" bson:"class_name"`
	Email     string `json:"email,omitempty" bson:"email"`
	Age       int    `json:"age,omitempty" bson:"age"`
}

type Student struct {
	ID        int    `json:"id,omitempty" bson:"id"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	ClassName string `json:"class_name,omitempty" bson:"class_name"`
	Email     string `json:"email,omitempty" bson:"email"`
	Age       int    `json:"age,omitempty" bson:"age"`
}

type StudentSearchRequest struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name"`
	LastName  string `json:"last_name,omitempty" bson:"last_name"`
	ClassName string `json:"class_name,omitempty" bson:"class_name"`
	Email     string `json:"email,omitempty" bson:"email"`
	Age       int    `json:"age" bson:"age"`
}

type UserClaims struct {
	UserID int    `json:"uid"`
	Phone  string `json:"p"`
	Email  string `json:"e"`
	jwt.StandardClaims
}

type Count struct {
	ID    int `bson:"id" json:"id"`
	Value int `bson:"val" json:"val"`
}

type Error struct {
	Code int
	Msg  string
}
