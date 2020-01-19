package route

import (
	"github.com/Splash07/nc_student/config"
	"github.com/Splash07/nc_student/handler"
	"github.com/Splash07/nc_student/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func All(e *echo.Echo) {
	Staff(e)
	Public(e)
}

func Staff(e *echo.Echo) {
	g := e.Group("/api/student/v1/staff")
	JWTConfig := middleware.JWTConfig{
		SigningKey: []byte(config.Config.JWTSecret.JWTKey),
		Claims:     &model.UserClaims{},
	}
	g.Use(middleware.JWTWithConfig(JWTConfig))
	g.POST("/student", handler.AddStudent)
	g.PUT("/student", handler.UpdateStudent)
	g.DELETE("/student", handler.DeleteStudentId)

}

func Public(e *echo.Echo) {
	g := e.Group("/api/student/v1/public")
	g.GET("/health", handler.HealthCheck)
	//g.POST("/test", handler.AddStudentAutoIncrement)
	g.GET("/all_students", handler.GetAllStudents)
	g.PATCH("/student_simple", handler.SearchStudentSimple)
	g.PATCH("/student_custom", handler.SearchStudentCustom)
}
