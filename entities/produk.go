package entities

import "gorm.io/gorm"

type Produk struct {
	gorm.Model

	Nama         string      `json:"nama"`
	Stok         int         `json:"stok"`
	Transaksi_id []Transaksi `json:"produk_id" gorm:"foreignKey:produk_id"`
}
