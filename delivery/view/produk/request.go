package produk

type InsertProdukRequest struct {
	Nama string `json:"nama" validate:"required"`
	Stok int    `json:"stok"`
}

// type UpdateStokRequest struct {
// 	Pengarang string `json:"pengarang" validate:"required"`
// }
