package entities

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	Produk_id       uint   `json:"produk_id"`
	User_id         uint   `json:"user_id"`
	Produk          string `json:"produk"`
	Qty             int    `json:"qty"`
	Jenis_transaksi string `json:"jenis_transaksi"`
}
