package produk

import "projek/be8/entities"

type Produk interface {
	InsertProduk(newProduk entities.Produk) (entities.Produk, error)
	GetAllProduk() ([]entities.Produk, error)
	GetProdukID(ID uint) (entities.Produk, error)
	UpdateProduk(produk *entities.Produk) (entities.Produk, error)
}
