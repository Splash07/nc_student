package route

import (
	"github.com/Splash07/nc_student/handler"
	"github.com/labstack/echo/v4"
)

func All(e *echo.Echo) {
	Private(e)
	Staff(e)
	Public(e)
}

func Private(e *echo.Echo) {

}

func Staff(e *echo.Echo) {
	g := e.Group("/api/student/v1/staff")
	g.POST("/student", handler.AddStudent)
	g.PUT("/student", handler.UpdateStudent)

}

func Public(e *echo.Echo) {
	g := e.Group("/api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	g.GET("/all_students", handler.GetAllStudents)
	g.PATCH("/student_simple", handler.SearchStudentSimple)

}
