package controller

import (
	vtransaksi "projek/be8/delivery/view/transaksi"
	"projek/be8/repository/produk"
	"projek/be8/repository/transaksi"

	"net/http"
	"projek/be8/entities"

	"github.com/golang-jwt/jwt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TransaksiController struct {
	Rproduk produk.Produk
	Repo    transaksi.Transaksi
	Valid   *validator.Validate
}

func New(repo transaksi.Transaksi, rproduk produk.Produk, valid *validator.Validate) *TransaksiController {
	return &TransaksiController{
		Rproduk: rproduk,
		Repo:    repo,
		Valid:   valid,
	}
}

func (tc *TransaksiController) InsertTransaksi(c echo.Context) error {
	var tmpTransaksi vtransaksi.InsertTransaksiRequest

	if err := c.Bind(&tmpTransaksi); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, "fail")
	}

	if err := tc.Valid.Struct(tmpTransaksi); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, "fail")
	}

	if tmpTransaksi.Jenis_transaksi != "pembelian" && tmpTransaksi.Jenis_transaksi != "penjualan" {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	newTransaksi := entities.Transaksi{User_id: uint(tmpTransaksi.User_id), Produk_id: tmpTransaksi.Produk_id, Produk: tmpTransaksi.Produk, Qty: tmpTransaksi.Qty, Jenis_transaksi: tmpTransaksi.Jenis_transaksi}
	res, err := tc.Repo.Insert(newTransaksi)
	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, "fail")
	}

	produk, _ := tc.Rproduk.GetProdukID(res.Produk_id)

	if newTransaksi.Jenis_transaksi == "pembelian" {
		produk.Stok += tmpTransaksi.Qty
	}

	if newTransaksi.Jenis_transaksi == "penjualan" {
		produk.Stok -= tmpTransaksi.Qty
	}

	_, err = tc.Rproduk.UpdateProduk(&produk)

	if err != nil {
		log.Warn("tidak bisa update stok")
		return c.JSON(http.StatusInternalServerError, "fail")
	}

	log.Info("success insert data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       http.StatusOK,
		"message":    "berhasil update data produk",
		"status":     true,
		"Total Stok": produk.Stok,
		"data":       res,
	})
}

func (tc *TransaksiController) GetAllTransaksi(c echo.Context) error {

	res, err := tc.Repo.GetAll()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, "fail")
	}
	log.Info("berhasil get all data")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"count":   len(res),
		"message": "berhasil get all data",
		"status":  true,
		"data":    res,
	})
}

func (pc *TransaksiController) GetTransaksi(c echo.Context) error {
	tipe := c.Param("tipe")

	hasil, err := pc.Repo.GetTrans(tipe)

	if err != nil {
		log.Warn(err)
		notFound := "data tidak ditemukan"
		if err.Error() == notFound {
			return c.JSON(http.StatusNotFound, "fail")
		}
		return c.JSON(http.StatusInternalServerError, "fail")

	}

	log.Info("data pembelian found")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"count":   len(hasil),
		"message": "riwayat transaksi ditemukan",
		"status":  true,
		"data":    hasil,
	})

}

func ExtractTokenUserId(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return userId
	}
	return 0
}
