package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerProduk interface {
	InsertProd(c echo.Context) error
	GetAllProd(c echo.Context) error
	GetProdID(c echo.Context) error
}
