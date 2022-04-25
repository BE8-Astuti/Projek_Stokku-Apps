package transaksi

import (
	"net/http"
	"projek/be8/entities"
)

func SuccessInsert(data entities.Transaksi) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "berhasil insert data book",
		"status":  true,
		"data":    data,
	}
}

func BadRequest() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terdapat kesalahan pada input data transaksi",
		"status":  false,
		"data":    nil,
	}
}

func SuccessUpdate(data entities.Transaksi) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusContinue,
		"message": "berhasil update data transaksi",
		"status":  true,
		"data":    data,
	}
}

func BadRequestUpdate() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "terjadi kesalahan update data transaksi",
		"status":  false,
		"data":    nil,
	}
}
