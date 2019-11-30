package handler

import (
	"github.com/Splash07/nc_student/db"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestDB(c echo.Context) error {
	db.Test()
	return c.String(http.StatusOK, "Test DB Sucessful!")
}
