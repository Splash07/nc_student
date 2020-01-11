package middleware

import (
	"github.com/labstack/echo/v4"
	"log"
)

func SimpleLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			resp := c.Response()
			next(c)

			log.Println(req.URL.Host, req.URL.Path, req.Method, resp.Status)

			return nil
		}
	}
}
