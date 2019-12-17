package db

type StudentUpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ClassName string `json:"class_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
}

type Error struct {
	Code int
	Msg  string
}
