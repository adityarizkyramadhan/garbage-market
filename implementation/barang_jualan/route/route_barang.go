package route

import (
	serviceBarang "github.com/adityarizkyramadhan/garbage-market/implementation/barang_jualan/db/postgres"
	"github.com/adityarizkyramadhan/garbage-market/implementation/barang_jualan/delivery/http"
	serviceToko "github.com/adityarizkyramadhan/garbage-market/implementation/toko_sampah/db/postgres"
	"github.com/adityarizkyramadhan/garbage-market/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouteBarang(r *gin.RouterGroup, db *gorm.DB) {
	service := serviceBarang.NewBarangJualanService(db)
	serviceToko := serviceToko.NewTokoSampahService(db)
	handler := http.NewHandlerBarangJualan(service, serviceToko)
	r.GET("/:id", middleware.ValidateJWToken(), handler.GetBarangJualanById)
	r.GET("/all", middleware.ValidateJWToken(), handler.GetBarangJualanAll)
	r.DELETE("/:id", middleware.ValidateJWToken(), handler.DeleteBarangJualan)
	r.POST("/add", middleware.ValidateJWToken(), handler.CreateBarangJualan)
}
