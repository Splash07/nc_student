package route

import (
	"github.com/labstack/echo/v4"
	"github.com/letanthang/nc_student/handler"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

func Staff(e *echo.Echo) {

}

func Public(e *echo.Echo) {
	g := e.Group("/api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
}
