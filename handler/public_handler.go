package handler

import (
	"net/http"

	"github.com/Splash07/nc_student/db"
	mw "github.com/Splash07/nc_student/middleware"
	"github.com/Splash07/nc_student/model"
	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func GetAllStudents(c echo.Context) error {
	students, err := db.GetAllStudent()
	if err != nil {
		mw.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, students)
}

func SearchStudentSimple(c echo.Context) error {
	var req model.StudentSearchRequest
	if err := c.Bind(&req); err != nil { // aslo acts as validation here
		mw.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	students, err := db.SearchStudentSimple(req)
	if err != nil {
		mw.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, students)
}

func SearchStudentCustom(c echo.Context) error {
	var req model.StudentSearchRequest
	if err := c.Bind(&req); err != nil { // aslo acts as validation here
		mw.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	students, err := db.SearchStudentCustom(req)
	if err != nil {
		mw.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, students)
}
