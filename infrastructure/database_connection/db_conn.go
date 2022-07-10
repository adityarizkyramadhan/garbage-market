package database_connection

import (
	"fmt"
	"github.com/adityarizkyramadhan/garbage-market/domain"

	"github.com/adityarizkyramadhan/garbage-market/infrastructure/database_driver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MakeConnection(data database_driver.DriverSupabase) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s "+
		"password=%s "+
		"host=%s "+
		"TimeZone=Asia/Singapore "+
		"port=%s "+
		"dbname=%s", data.User, data.Password, data.Host, data.Port, data.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.MetaUser{}, &domain.User{}, domain.TokoSampah{}, &domain.BarangJualan{}, &domain.BasketJualan{}, &domain.PembayaranUser{}); err != nil {
		return nil, err
	}
	return db, nil
}
