package middleware

import (
	"time"

	"github.com/Splash07/nc_student/db"
	"github.com/Splash07/nc_student/model"
	"github.com/Splash07/nc_student/utils"
	"github.com/labstack/echo/v4"
)

func LoggerFile() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			resp := c.Response()
			next(c)

			utils.GeneralLogger.Println(req.URL.Path, req.Method, resp.Status)
			return nil
		}
	}
}

func LoggerMongoDB() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			resp := c.Response()
			next(c)

			go func() {
				var log model.Log
				log.Date = time.Now()
				log.Method = req.Method
				log.Path = req.URL.Path
				log.Status = resp.Status

				db.AddLog(&log)
			}()
			return nil
		}
	}
}
