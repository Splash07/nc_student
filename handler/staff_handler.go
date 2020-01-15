package handler

import (
	"net/http"

	"github.com/Splash07/nc_student/db"
	mw "github.com/Splash07/nc_student/middleware"
	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	var student db.Student
	if err := c.Bind(&student); err != nil {
		mw.ErrorLogger.Printf("Binding error occured: %v", err)
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddStudent(&student)
	if err != nil {
		mw.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func AddStudentAutoIncrement(c echo.Context) {
	// var student db.Student
	// if err := c.Bind(&student); err != nil {
	// 	return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	// }
	db.AddStudentAutoIncrement()
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	// }

}

func UpdateStudent(c echo.Context) error {

	var student db.StudentUpdateRequest
	if err := c.Bind(&student); err != nil {
		mw.ErrorLogger.Printf("Binding error occured: %v", err)
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateStudent(&student)
	if err != nil {
		mw.ErrorLogger.Printf("Update error occured: %v", err)
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, res)
}
