package system

import (
	"bookstore/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	echo *echo.Echo
	controller controllers.Controller
}

func (server *Server) Init() {

	dsn := "root:@tcp(127.0.0.1:3306)/golang_bootcamp?charset=utf8mb4&parseTime=True&loc=Local"
	db, err :=  gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	server.echo = echo.New()
	server.echo.Use(middleware.Logger())
	server.echo.Use(middleware.Recover())
	server.controller.DeclareModel(db)
	server.controller.DeclareRoutes(server.echo)
}

func (server *Server) Start() {
	server.echo.Logger.Fatal(server.echo.Start(":1323"))
}