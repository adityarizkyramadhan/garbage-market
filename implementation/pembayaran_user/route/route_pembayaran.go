package route

import (
	"github.com/adityarizkyramadhan/garbage-market/implementation/pembayaran_user/db/postgres"
	"github.com/adityarizkyramadhan/garbage-market/implementation/pembayaran_user/delivery/http"
	"github.com/adityarizkyramadhan/garbage-market/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoutePembayaran(r *gin.RouterGroup, db *gorm.DB) {
	service := postgres.NewPembayaranService(db)
	handler := http.NewDeliveryPembayaran(service)
	r.POST("/:id", middleware.ValidateJWToken(), handler.CreatePembayaran)
	r.GET("/:id", middleware.ValidateJWToken(), handler.GetPembayaranById)
	r.GET("/user/:id", middleware.ValidateJWToken(), handler.GetPembayaranByIdUser)
}
