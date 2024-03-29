package http

import (
	"net/http"
	"strconv"

	"github.com/adityarizkyramadhan/garbage-market/domain"
	"github.com/adityarizkyramadhan/garbage-market/middleware"
	"github.com/adityarizkyramadhan/garbage-market/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type deliveryMetadata struct {
	pg domain.ServiceMetaData
}

func NewDeliveryMetadata(pg domain.ServiceMetaData) domain.HandlerMetaData {
	return &deliveryMetadata{
		pg: pg,
	}
}

type inputLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *deliveryMetadata) Login(c *gin.Context) {
	var input inputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data, err := d.pg.Login(input.Email)
	if err != nil || data == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), data))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Login success", gin.H{
		"token": token,
	}))
}

type inputRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Pin      string `json:"pin"`
}

func (d *deliveryMetadata) Register(c *gin.Context) {
	var input inputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data, err := d.pg.Register(&domain.MetaUser{
		Email:    input.Email,
		Password: string(hashedPass),
	})
	if err != nil || data == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), data))
	}
	token, err := middleware.GenerateJWToken(data.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, utils.ResponseWhenSuccess("Login success", gin.H{
		"token": token,
	}))
}

func (d *deliveryMetadata) GetUserById(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), nil))
		return
	}
	data, err := d.pg.GetUserById(uint(idInt))
	if err != nil || data == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), data))
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Get user by id success", data))
}

func (d *deliveryMetadata) GetProfile(c *gin.Context) {
	idUSer := uint(c.MustGet("id").(float64))
	data, err := d.pg.GetUserById(idUSer)
	if err != nil || data == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error(), data))
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("Get profile success", data))
}
