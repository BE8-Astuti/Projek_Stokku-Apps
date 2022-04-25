package user

import (
	"net/http"
	"projek/be8/entities"
)

type LoginResponse struct {
	Data  entities.User
	Token string
}

func SuccessInsert(data entities.User) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data user",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data user",
		"status":  false,
		"data":    nil,
	}
}
