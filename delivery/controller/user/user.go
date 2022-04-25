package controller

import (
	"projek/be8/delivery/view"
	userview "projek/be8/delivery/view/user"
	ruser "projek/be8/repository/user"
	"time"

	"net/http"
	"projek/be8/entities"

	"github.com/golang-jwt/jwt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct {
	Repo  ruser.User
	Valid *validator.Validate
}

func New(repo ruser.User, valid *validator.Validate) *UserController {
	return &UserController{
		Repo:  repo,
		Valid: valid,
	}
}

func (uc *UserController) InsertUser(c echo.Context) error {
	var tmpUser userview.InsertUserRequest

	if err := c.Bind(&tmpUser); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	if err := uc.Valid.Struct(tmpUser); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	newUser := entities.User{Name: tmpUser.Name, Password: tmpUser.Password, Phone: tmpUser.Phone}
	res, err := uc.Repo.InsertUser(newUser)

	if err != nil {
		log.Warn("masalah pada server")
		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
	}
	log.Info("berhasil insert")
	return c.JSON(http.StatusCreated, userview.SuccessInsert(res))
}

// func (uc *UserController) GetAllUser(c echo.Context) error {

// 	res, err := uc.Repo.GetAllUser()

// 	if err != nil {
// 		log.Warn("masalah pada server")
// 		return c.JSON(http.StatusInternalServerError, view.InternalServerError())
// 	}
// 	log.Info("berhasil get all data")
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"code":    http.StatusOK,
// 		"message": "berhasil get all data",
// 		"status":  true,
// 		"data":    res,
// 	})
// }

func (uc *UserController) Login(c echo.Context) error {
	param := userview.LoginRequest{}

	if err := c.Bind(&param); err != nil {
		log.Warn("salah input")
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	if err := uc.Valid.Struct(param); err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, userview.BadRequest())
	}

	hasil, err := uc.Repo.Login(param.Name, param.Password)

	if err != nil {
		log.Warn(err.Error())
		return c.JSON(http.StatusNotFound, "HP atau Password tidak ditemukan")
	}

	res := userview.LoginResponse{Data: hasil}

	if res.Token == "" {
		token, _ := CreateToken(int(hasil.ID))
		res.Token = token
		return c.JSON(http.StatusOK, view.OK(res, "Berhasil login"))
	}

	return c.JSON(http.StatusOK, view.OK(res, "Berhasil login"))
}

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["expired"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("A$T0T!"))
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
