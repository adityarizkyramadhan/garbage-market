package route

import (
	"github.com/adityarizkyramadhan/garbage-market/implementation/basket_jualan/db/postgres"
	"github.com/adityarizkyramadhan/garbage-market/implementation/basket_jualan/delivery/http"
	"github.com/adityarizkyramadhan/garbage-market/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteBasket(r *gin.RouterGroup, db *gorm.DB) {
	service := postgres.NewbasketService(db)
	handler := http.NewHandlerBasket(service)
	r.POST("/newbasket", middleware.ValidateJWToken(), handler.CreateBasketJualan)
	r.POST("/updatebasket", middleware.ValidateJWToken(), handler.UpdateBasketJualan)
	r.GET("/basket/:id", middleware.ValidateJWToken(), handler.GetBasketJualanById)
	r.GET("/basket", middleware.ValidateJWToken(), handler.GetAllBasketJualanByUser)
}
