package transaksi

type InsertTransaksiRequest struct {
	User_id         uint   `json:"user_id"`
	Produk_id       uint   `json:"produk_id"`
	Produk          string `json:"produk" validate:"required"`
	Qty             int    `json:"qty"`
	Jenis_transaksi string `json:"jenis_transaksi"`
}
