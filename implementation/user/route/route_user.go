package route

import (
	"github.com/adityarizkyramadhan/garbage-market/implementation/user/db/postgres"
	"github.com/adityarizkyramadhan/garbage-market/implementation/user/delivery/http"
	"github.com/adityarizkyramadhan/garbage-market/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteMetaData(r gin.RouterGroup, db *gorm.DB) {
	service := postgres.NewUserService(db)
	handler := http.NewDeliveryUser(service)
	r.POST("/newuser", middleware.ValidateJWToken(), handler.CreateUser)
	r.POST("/update", middleware.ValidateJWToken(), handler.UpdateUser)
	r.POST("/user/:id", middleware.ValidateJWToken(), handler.GetUserById)
}
