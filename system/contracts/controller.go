package contracts

import "github.com/labstack/echo/v4"

type Controller interface {
	Route(context echo.Context)
	Index(context echo.Context)
	Create(context echo.Context)
	Show(context echo.Context)
	Update(context echo.Context)
	Delete(context echo.Context)
}