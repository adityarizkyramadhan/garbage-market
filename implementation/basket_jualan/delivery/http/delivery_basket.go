package http

import (
	"net/http"
	"strconv"

	"github.com/adityarizkyramadhan/garbage-market/domain"
	"github.com/adityarizkyramadhan/garbage-market/utils"
	"github.com/gin-gonic/gin"
)

type handlerBasket struct {
	service domain.BasketJualanService
}

func NewHandlerBasket(service domain.BasketJualanService) domain.DeliveryBasket {
	return &handlerBasket{service}
}

type input struct {
	JumlahBarang int `json:"jumlah_barang"`
}

//Create new basket jualan
func (b *handlerBasket) CreateBasketJualan(c *gin.Context) {
	var basket input
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	idUser := c.MustGet("id").(uint)
	idBarang := c.Param("id")
	id, err := strconv.Atoi(idBarang)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data := &domain.BasketJualan{
		JumlahBarang:   basket.JumlahBarang,
		IdMetaUser:     idUser,
		IdBarangJualan: id,
	}
	if err := b.service.CreateBasketJualan(data); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("Basket jualan berhasil dibuat", basket))
}

//Get all basket jualan by user
func (b *handlerBasket) GetAllBasketJualanByUser(c *gin.Context) {
	idUser := c.MustGet("id").(int)
	data, err := b.service.GetBasketJualanByIdUser(uint(idUser))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil mengambil data", data))
}

//get baseket jualan by id
func (b *handlerBasket) GetBasketJualanById(c *gin.Context) {
	id := c.Param("id")
	idBusket, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data, err := b.service.GetBasketJualanById(uint(idBusket))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil mengambil data", data))
}

//delete basket jualan
func (b *handlerBasket) DeleteBasketJualan(c *gin.Context) {
	id := c.Param("id")
	idBusket, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	if err := b.service.DeleteBasketJualan(uint(idBusket)); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil menghapus data", nil))
}

//update basket jualan
func (b *handlerBasket) UpdateBasketJualan(c *gin.Context) {
	var basket input
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	id := c.Param("id")
	idBusket, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data := &domain.BasketJualan{
		JumlahBarang:   basket.JumlahBarang,
		IdMetaUser:     c.MustGet("id").(uint),
		IdBarangJualan: idBusket,
	}
	if err := b.service.UpdateBasketJualan(uint(idBusket), data); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil mengubah data", nil))
}
