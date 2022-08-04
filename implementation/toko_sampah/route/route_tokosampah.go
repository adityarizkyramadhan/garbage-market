package route

import (
	"github.com/adityarizkyramadhan/garbage-market/implementation/toko_sampah/db/postgres"
	"github.com/adityarizkyramadhan/garbage-market/implementation/toko_sampah/delivery/http"
	"github.com/adityarizkyramadhan/garbage-market/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteTokoSampah(r *gin.RouterGroup, db *gorm.DB) {
	service := postgres.NewTokoSampahService(db)
	handler := http.NewDeliveryTokoSampah(service)
	r.POST("/newtokosampah", middleware.ValidateJWToken(), handler.CreateTokoSampah)
	r.POST("/updatetokosampah", middleware.ValidateJWToken(), handler.UpdateTokoSampah)
	r.GET("/single/:id", middleware.ValidateJWToken(), handler.GetTokoSampahById)
	r.GET("/user", middleware.ValidateJWToken(), handler.GetTokoByIdUser)
	r.GET("/all", middleware.ValidateJWToken(), handler.GetAllTokoSampah)
}
