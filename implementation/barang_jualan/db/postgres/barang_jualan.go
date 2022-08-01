package postgres

import (
	"github.com/adityarizkyramadhan/garbage-market/domain"
	"gorm.io/gorm"
)

type barangJualanService struct {
	db *gorm.DB
}

func NewBarangJualanService(db *gorm.DB) domain.BarangJualanService {
	return &barangJualanService{db}
}

//create new barang jualan
func (b *barangJualanService) CreateBarangJualan(barang *domain.BarangJualan) error {
	return b.db.Create(barang).Error
}

func (b *barangJualanService) GetBarangJualanById(id uint) (*domain.BarangJualan, error) {
	barang := &domain.BarangJualan{}
	err := b.db.First(barang, id).Error
	return barang, err
}

func (b *barangJualanService) UpdateBarangJualan(barang *domain.BarangJualan) error {
	return b.db.Save(barang).Error
}

func (b *barangJualanService) DeleteBarangJualan(id uint) error {
	barang := &domain.BarangJualan{}
	return b.db.Delete(barang, id).Error
}

func (b *barangJualanService) GetBarangJualanAll() ([]*domain.BarangJualan, error) {
	barang := []*domain.BarangJualan{}
	err := b.db.Find(&barang).Error
	return barang, err
}
