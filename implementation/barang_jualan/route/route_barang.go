package route

import (
	"github.com/adityarizkyramadhan/garbage-market/implementation/barang_jualan/db/postgres"
	"github.com/adityarizkyramadhan/garbage-market/implementation/barang_jualan/delivery/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouteBarang(r *gin.RouterGroup, db *gorm.DB) {
	service := postgres.NewBarangJualanService(db)
	handler := http.NewHandlerBarangJualan(service)
	r.GET("/:id", handler.GetBarangJualanById)
	r.POST("/", handler.CreateBarangJualan)
}
