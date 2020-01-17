package handler

import (
	"net/http"

	"github.com/Splash07/nc_student/db"
	mw "github.com/Splash07/nc_student/middleware"
	"github.com/Splash07/nc_student/model"
	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		mw.ErrorLogger.Printf("Binding error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddStudent(&student)
	if err != nil {
		mw.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateStudent(c echo.Context) error {

	var student model.StudentUpdateRequest
	if err := c.Bind(&student); err != nil {
		mw.ErrorLogger.Printf("Binding error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateStudent(&student)
	if err != nil {
		mw.ErrorLogger.Printf("Update error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, res)
}
