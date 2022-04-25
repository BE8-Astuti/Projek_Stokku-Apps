package controller

import (
	"projek/be8/delivery/view"
	vproduk "projek/be8/delivery/view/produk"
	"projek/be8/repository/produk"
	"strconv"

	"net/http"
	"projek/be8/entities"

	"github.com/golang-jwt/jwt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type ProdukController struct {
	Repo  produk.Produk
	Valid *validator.Validate
}

func New(repo produk.Produk, valid *validator.Validate) *ProdukController {
	return &ProdukController{
		Repo:  repo,
		Valid: valid,
	}
}

func (pc *ProdukController) InsertProd(c echo.Context) error {
	var tmpProd vproduk.InsertProdukRequest

	if err := c.Bind(&tmpProd); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, "fail")
	}

	if err := pc.Valid.Struct(tmpProd); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, "fail")
	}

	newProd := entities.Produk{Nama: tmpProd.Nama, Stok: tmpProd.Stok}
	res, err := pc.Repo.InsertProduk(newProd)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, "fail")
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, res)
}

func (pc *ProdukController) GetAllProd(c echo.Context) error {

	res, err := pc.Repo.GetAllProduk()

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, "fail")
	}
	log.Info("berhasil get all data produk")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"count":   len(res),
		"message": "berhasil get all data produk",
		"status":  true,
		"data":    res,
	})
}

func (pc *ProdukController) GetProdID(c echo.Context) error {
	id := c.Param("id")

	convID, err := strconv.Atoi(id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "connot convert ID",
			"data":    nil,
		})
	}

	hasil, err := pc.Repo.GetProdukID(uint(convID))

	if err != nil {
		log.Warn(err)
		notFound := "data tidak ditemukan"
		if err.Error() == notFound {
			return c.JSON(http.StatusNotFound, view.NotFoundError())
		}
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())

	}

	log.Info("data produk found")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "data produk ditemukan",
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
