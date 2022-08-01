package domain

import "github.com/gin-gonic/gin"

type ServiceMetaData interface {
	Register(user *MetaUser) (*MetaUser, error)
	Login(email string) (*MetaUser, error)
	GetUserById(id uint) (*MetaUser, error)
}

type HandlerMetaData interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetUserById(c *gin.Context)
}

type ServiceUser interface {
	GetUserById(id uint) (*User, error)
	UpdateUser(id uint, user *User) error
	CreateUser(user *User) error
}

type HandlerUser interface {
	GetUserById(c *gin.Context)
	UpdateUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type TokoSampahService interface {
	CreateTokoSampah(toko *TokoSampah) error
	GetTokoSampahById(id uint) (*TokoSampah, error)
	UpdateTokoSampah(toko *TokoSampah) error
	GetTokoByIdUser(idUser uint) (*TokoSampah, error)
}

type HandlerTokoSampah interface {
	CreateTokoSampah(c *gin.Context)
	GetTokoSampahById(c *gin.Context)
	UpdateTokoSampah(c *gin.Context)
	GetTokoByIdUser(c *gin.Context)
}

type BasketJualanService interface {
	CreateBasketJualan(basket *BasketJualan) error
	GetBasketJualanById(id uint) (*BasketJualan, error)
	DeleteBasketJualan(idBusket uint) error
	GetBasketJualanByIdUser(idUser uint) (*[]BasketJualan, error)
	UpdateBasketJualan(id uint, basket *BasketJualan) error
}

type DeliveryBasket interface {
	CreateBasketJualan(c *gin.Context)
	GetAllBasketJualanByUser(c *gin.Context)
	GetBasketJualanById(c *gin.Context)
	DeleteBasketJualan(c *gin.Context)
	UpdateBasketJualan(c *gin.Context)
}

type BarangJualanService interface {
	CreateBarangJualan(barang *BarangJualan) error
	GetBarangJualanById(id uint) (*BarangJualan, error)
	GetBarangJualanAll() ([]*BarangJualan, error)
	UpdateBarangJualan(barang *BarangJualan) error
	DeleteBarangJualan(id uint) error
}

type HandlerBarangJualan interface {
	CreateBarangJualan(c *gin.Context)
	GetBarangJualanById(c *gin.Context)
	GetBarangJualanAll(c *gin.Context)
	DeleteBarangJualan(c *gin.Context)
}

type PembayaranService interface {
	CreatePembayaran(pembayaran *PembayaranUser) error
	GetPembayaranById(id uint) (*PembayaranUser, error)
	GetPembayaranByIdUser(id uint) (*[]PembayaranUser, error)
}
