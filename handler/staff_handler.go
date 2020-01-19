package handler

import (
	"net/http"

	"github.com/Splash07/nc_student/db"
	"github.com/Splash07/nc_student/model"
	"github.com/Splash07/nc_student/utils"
	"github.com/labstack/echo/v4"
)

func AddStudent(c echo.Context) error {
	var student model.Student
	if err := c.Bind(&student); err != nil {
		utils.ErrorLogger.Printf("Binding error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.AddStudent(&student)
	if err != nil {
		utils.ErrorLogger.Printf("Error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, res)
}

func UpdateStudent(c echo.Context) error {

	var student model.StudentUpdateRequest
	if err := c.Bind(&student); err != nil {
		utils.ErrorLogger.Printf("Binding error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.UpdateStudent(&student)
	if err != nil {
		utils.ErrorLogger.Printf("Update error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteStudentId(c echo.Context) error {
	var student model.StudentDeleteRequest
	if err := c.Bind(&student); err != nil {
		utils.ErrorLogger.Printf("Binding error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	res, err := db.DeleteStudentById(&student)
	if err != nil {
		utils.ErrorLogger.Printf("Delete error occured: %v", err)
		return c.JSON(http.StatusBadRequest, model.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}
	return c.JSON(http.StatusOK, res)
}
