package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/adityarizkyramadhan/garbage-market/domain"
	"github.com/adityarizkyramadhan/garbage-market/utils"
	"github.com/gin-gonic/gin"
	storage_go "github.com/supabase-community/storage-go"
)

type deliveryUser struct {
	pg domain.ServiceUser
}

func NewDeliveryUser(pg domain.ServiceUser) domain.HandlerUser {
	return &deliveryUser{
		pg: pg,
	}
}

func (d *deliveryUser) CreateUser(c *gin.Context) {
	id := c.MustGet("id").(float64)
	fmt.Println(id)
	nama := c.Request.FormValue("nama")
	tanggalLahir := c.Request.FormValue("tanggal_lahir")
	gender := c.Request.FormValue("gender")
	alamat := c.Request.FormValue("alamat")
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
	fileName := fmt.Sprintf("profile%s%s%s", nama, fileInput.Filename, time.Now().Format("20060102150405"))
	fileName = strings.ReplaceAll(fileName, " ", "")
	client.UploadFile("foto-proker", fileName, file)
	linkImage := utils.GenerateLinkImage(fileName)
	user := domain.User{
		Nama:         nama,
		TanggalLahir: tanggalLahir,
		Gender:       gender,
		Alamat:       alamat,
		LinkFoto:     linkImage,
		IdMetaUser:   uint(id),
	}
	if err := d.pg.CreateUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseWhenFail("Fail to save user", nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Success to save user", user))
}

func (d *deliveryUser) GetUserById(c *gin.Context) {
	id := c.MustGet("id").(uint)
	user, err := d.pg.GetUserById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseWhenFail("Fail to get user", err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Success to get user", user))
}

func (d *deliveryUser) UpdateUser(c *gin.Context) {
	id := c.MustGet("id").(uint)
	nama := c.Request.FormValue("nama")
	tanggalLahir := c.Request.FormValue("tanggal_lahir")
	gender := c.Request.FormValue("gender")
	alamat := c.Request.FormValue("alamat")
	user := &domain.User{
		Nama:         nama,
		TanggalLahir: tanggalLahir,
		Gender:       gender,
		Alamat:       alamat,
	}
	if err := d.pg.UpdateUser(id, user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseWhenFail("Fail to update user", nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Success to update user", user))
}
