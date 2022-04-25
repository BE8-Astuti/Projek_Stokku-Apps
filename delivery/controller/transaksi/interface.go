package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerTransaksi interface {
	InsertTransaksi(c echo.Context) error
	// InsertTransaksiJual(c echo.Context) error
	GetAllTransaksi(c echo.Context) error
	GetTransaksi(c echo.Context) error
	// UpdateBookID(c echo.Context) error
	// DeleteBookID(c echo.Context) error
}
