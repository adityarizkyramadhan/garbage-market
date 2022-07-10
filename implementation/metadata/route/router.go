package route

import (
	_dbMetadata "github.com/adityarizkyramadhan/garbage-market/implementation/metadata/db/postgres"
	_handlerMetadata "github.com/adityarizkyramadhan/garbage-market/implementation/metadata/delivery/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RouteMetaData(r gin.RouterGroup, db *gorm.DB) {
	service := _dbMetadata.NewMetaDataService(db)
	handler := _handlerMetadata.NewDeliveryMetadata(service)
	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
}
