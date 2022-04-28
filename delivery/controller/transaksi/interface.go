package controller

import (
	"github.com/labstack/echo/v4"
)

type ControllerTransaksi interface {
	InsertTransaksi(c echo.Context) error

	GetAllTransaksi(c echo.Context) error
	GetTransaksi(c echo.Context) error
	RiwayatAllTrans(c echo.Context) error
	HistoriTrans(c echo.Context) error
}
