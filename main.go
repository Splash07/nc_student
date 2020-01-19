package main

import (
	"fmt"
	"log"

	"github.com/Splash07/nc_student/config"
	mw "github.com/Splash07/nc_student/middleware"
	"github.com/Splash07/nc_student/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("config app: %+v", config.Config)
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(mw.LoggerFile())
	e.Use(mw.LoggerMongoDB())
	route.All(e)

	log.Println(e.Start(":1010"))
}
