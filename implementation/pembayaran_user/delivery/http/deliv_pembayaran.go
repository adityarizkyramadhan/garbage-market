package http

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/adityarizkyramadhan/garbage-market/domain"
	"github.com/adityarizkyramadhan/garbage-market/utils"
	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
)

type deliveryPembayaran struct {
	service domain.PembayaranService
}

func NewDeliveryPembayaran(service domain.PembayaranService) *deliveryPembayaran {
	return &deliveryPembayaran{service}
}

func (d *deliveryPembayaran) CreatePembayaran(c *gin.Context) {
	idUser := c.MustGet("id").(int)
	idBarang := c.Param("id")
	id, err := strconv.Atoi(idBarang)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	jumlahPembelian := c.Request.FormValue("jumlah_pembelian")
	jumlahPembelianInt, _ := strconv.Atoi(jumlahPembelian)
	jumlahPembayaran := c.Request.FormValue("jumlah_pembayaran")
	jumlahPembayaranInt, _ := strconv.Atoi(jumlahPembayaran)
	fileInput, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail("Error when get file", err.Error()))
		return
	}
	file, err := fileInput.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail("Error when open file", err.Error()))
		return
	}
	client := storage_go.NewClient("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImpnanlqdnlsZG9hbXFuZGF6aXhsIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NDc4MzQ0MDQsImV4cCI6MTk2MzQxMDQwNH0.WVMjJIRoK_cnyfRXdYvTokNWBCCqLWfbeu7xXeZrs6I", nil)
	fileName := fmt.Sprintf("pembayaran%s%s", fileInput.Filename, time.Now().Format("20060102150405"))
	fileName = strings.ReplaceAll(fileName, " ", "")
	client.UploadFile("foto-proker", fileName, file)
	linkImage := utils.GenerateLinkImage(fileName)
	pembayaran := &domain.PembayaranUser{
		IdUser:           idUser,
		IdBarangJualan:   id,
		JumlahPembelian:  jumlahPembelianInt,
		JumlahPembayaran: jumlahPembayaranInt,
		LinkFoto:         linkImage,
	}
	if err := d.service.CreatePembayaran(pembayaran); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail("Error when create pembayaran", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("Success create pembayaran", pembayaran))

}

func (d *deliveryPembayaran) GetPembayaranById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	pembayaran, err := d.service.GetPembayaranById(uint(idInt))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseWhenFail("Error when get pembayaran", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Success get pembayaran", pembayaran))
}

func (d *deliveryPembayaran) GetPembayaranByIdUser(c *gin.Context) {
	id := c.MustGet("id").(int)
	pembayaran, err := d.service.GetPembayaranByIdUser(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, utils.ResponseWhenFail("Error when get pembayaran", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Success get pembayaran", pembayaran))
}
