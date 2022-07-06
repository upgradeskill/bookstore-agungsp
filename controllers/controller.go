package controllers

import (
	"bookstore/controllers/authcontroller"
	"bookstore/controllers/bookcontroller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller struct {
	authcontroller authcontroller.AuthController
	bookController bookcontroller.BookController
	/* Some Controller */
	/* Some Controller */
}

func (controller *Controller) DeclareModel(db *gorm.DB)  {
	controller.bookController.Model(db)
}

func (controller *Controller) DeclareRoutes(echo *echo.Echo) {
	controller.authcontroller.Route(echo)
	controller.bookController.Route(echo)
	/*
		controller.[someController].Route(echo)
		...
		...
	*/
}