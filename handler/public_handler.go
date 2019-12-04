package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Splash07/nc_student/db"
	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestDB(c echo.Context) error {
	db.Test()
	return c.String(http.StatusOK, "Test DB Sucessful!")
}

func GetAllStudents(c echo.Context) error {
	type Student struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Age       int    `json:"age"`
		ClassName string `json:"class_name"`
		Email     string `json:"email"`
	}
	inputJson := `[
		{"first_name":"Dylan", "last_name":"Nguyen", "age":20, "class_name":"Golang", "email":"nguyendat293@gmail.com"},
		{"first_name":"Dat", "last_name":"Nguyen", "age":24, "class_name":"Java", "email":"alexanderthegreat293@gmail.com"},
		{"first_name":"Ha", "last_name":"Hoang", "age":18, "class_name":"Front-end", "email":"hhh@gmail.com"},
		{"first_name":"Vu", "last_name":"Nguyen", "age":16, "class_name":"NodeJS", "email":"vu_nguyen@gmail.com"}
		]`

	var students []Student
	json.Unmarshal([]byte(inputJson), &students)
	return c.JSON(http.StatusOK, students)
}

func AddStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, db.Test)
}
