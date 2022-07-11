package route

import (
	routeBarang "github.com/adityarizkyramadhan/garbage-market/implementation/barang_jualan/route"
	basketRoute "github.com/adityarizkyramadhan/garbage-market/implementation/basket_jualan/route"
	metadataRoute "github.com/adityarizkyramadhan/garbage-market/implementation/metadata/route"
	pembayaranRoute "github.com/adityarizkyramadhan/garbage-market/implementation/pembayaran_user/route"
	tokoSampahRoute "github.com/adityarizkyramadhan/garbage-market/implementation/toko_sampah/route"
	userRoute "github.com/adityarizkyramadhan/garbage-market/implementation/user/route"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoute(r *gin.Engine, db *gorm.DB) {
	metadata := r.Group("/metadata")
	metadataRoute.RouteMetaData(metadata, db)
	user := r.Group("/user")
	userRoute.RouteUser(user, db)
	tokoSampah := r.Group("/toko_sampah")
	tokoSampahRoute.RouteTokoSampah(tokoSampah, db)
	pembayaran := r.Group("/pembayaran")
	pembayaranRoute.NewRoutePembayaran(pembayaran, db)
	barang := r.Group("/barang")
	routeBarang.NewRouteBarang(barang, db)
	basket := r.Group("/basket")
	basketRoute.RouteBasket(basket, db)

}
