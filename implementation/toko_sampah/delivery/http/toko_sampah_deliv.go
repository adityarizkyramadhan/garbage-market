package http

import (
	"net/http"
	"strconv"

	"github.com/adityarizkyramadhan/garbage-market/domain"
	"github.com/adityarizkyramadhan/garbage-market/utils"
	"github.com/gin-gonic/gin"
)

type deliveryTokoSampah struct {
	TokoSampahService domain.TokoSampahService
}

func NewDeliveryTokoSampah(TokoSampahService domain.TokoSampahService) domain.HandlerTokoSampah {
	return &deliveryTokoSampah{TokoSampahService}
}

type inputTokoSampah struct {
	NamaToko   string `json:"nama_toko"`
	AlamatToko string `json:"alamat_toko"`
}

func (d *deliveryTokoSampah) CreateTokoSampah(c *gin.Context) {
	idUser := c.MustGet("id").(float64)
	var input inputTokoSampah
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data := &domain.TokoSampah{
		NamaToko:   input.NamaToko,
		AlamatToko: input.AlamatToko,
		IdMetaUser: uint(idUser),
	}
	if err := d.TokoSampahService.CreateTokoSampah(data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil menambahkan toko sampah", data))
}

func (d *deliveryTokoSampah) GetTokoSampahById(c *gin.Context) {
	idToko := c.Param("id")
	id, err := strconv.Atoi(idToko)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data, err := d.TokoSampahService.GetTokoSampahById(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil mengambil data toko sampah", data))
}

func (d *deliveryTokoSampah) UpdateTokoSampah(c *gin.Context) {
	idUser := int(c.MustGet("id").(float64))
	idToko := c.Param("id")
	id, err := strconv.Atoi(idToko)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	var input inputTokoSampah
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data, err := d.TokoSampahService.GetTokoSampahById(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	if int(data.IdMetaUser) != idUser {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail("Tidak dapat mengakses data toko sampah ini", nil))
		return
	}
	data.NamaToko = input.NamaToko
	data.AlamatToko = input.AlamatToko
	if err := d.TokoSampahService.UpdateTokoSampah(data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil mengubah data toko sampah", data))
}

func (d *deliveryTokoSampah) GetTokoByIdUser(c *gin.Context) {
	idUser := int(c.MustGet("id").(float64))
	data, err := d.TokoSampahService.GetTokoByIdUser(uint(idUser))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil mengambil data toko sampah", data))
}

func (d *deliveryTokoSampah) GetAllTokoSampah(c *gin.Context) {
	data, err := d.TokoSampahService.GetAllTokoSampah()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Berhasil mengambil data toko sampah", data))
}
