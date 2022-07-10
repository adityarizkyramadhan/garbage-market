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

type handlerBarangJualan struct {
	service domain.BarangJualanService
}

func NewHandlerBarangJualan(service domain.BarangJualanService) domain.HandlerBarangJualan {
	return &handlerBarangJualan{service}
}

func (h *handlerBarangJualan) CreateBarangJualan(c *gin.Context) {
	namaBarang := c.Request.FormValue("nama_barang")
	hargaBarang := c.Request.FormValue("harga_barang")
	stokBarang := c.Request.FormValue("stok_barang")
	tipeBarang := c.Request.FormValue("tipe_barang")
	deskripsi := c.Request.FormValue("deskripsi")
	hargaBarangInt, _ := strconv.Atoi(hargaBarang)
	stokBarangInt, _ := strconv.Atoi(stokBarang)
	fileInput, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail("Error when get file", err.Error()))
		return
	}
	file, err := fileInput.Open()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail("Error when open file", err.Error()))
		return
	}
	client := storage_go.NewClient("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImpnanlqdnlsZG9hbXFuZGF6aXhsIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NDc4MzQ0MDQsImV4cCI6MTk2MzQxMDQwNH0.WVMjJIRoK_cnyfRXdYvTokNWBCCqLWfbeu7xXeZrs6I", nil)
	fileName := fmt.Sprintf("profile%s%s%s", namaBarang, fileInput.Filename, time.Now().Format("20060102150405"))
	fileName = strings.ReplaceAll(fileName, " ", "")
	client.UploadFile("foto-proker", fileName, file)
	linkImage := utils.GenerateLinkImage(fileName)
	barang := &domain.BarangJualan{
		NamaBarang:  namaBarang,
		HargaBarang: hargaBarangInt,
		StokBarang:  stokBarangInt,
		TipeBarang:  tipeBarang,
		Deskripsi:   deskripsi,
		LinkFoto:    linkImage,
	}
	if err := h.service.CreateBarangJualan(barang); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail("Error when create barang jualan", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("Success create barang jualan", barang))
}

func (h *handlerBarangJualan) GetBarangJualanById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	barang, err := h.service.GetBarangJualanById(uint(idInt))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail("Error when get barang jualan", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Success get barang jualan", barang))
}
