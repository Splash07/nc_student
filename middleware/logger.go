package middleware

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

var GeneralLogger *log.Logger
var ErrorLogger *log.Logger

func init() {

	absPath, err := filepath.Abs("D:/Study/github.com/Splash07/nc_student/log")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	GeneralLogger = log.New(generalLog, "General:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			resp := c.Response()
			next(c)

			GeneralLogger.Println(req.URL.Path, req.Method, resp.Status)
			return nil
		}
	}
}
