package postgres

import (
	"github.com/adityarizkyramadhan/garbage-market/domain"
	"gorm.io/gorm"
)

type pembayaranService struct {
	db *gorm.DB
}

func NewPembayaranService(db *gorm.DB) domain.PembayaranService {
	return &pembayaranService{db}
}

//create new pembayaran
func (p *pembayaranService) CreatePembayaran(pembayaran *domain.PembayaranUser) error {
	return p.db.Create(pembayaran).Error
}

//get pembayaran by id
func (p *pembayaranService) GetPembayaranById(id uint) (*domain.PembayaranUser, error) {
	pembayaran := &domain.PembayaranUser{}
	err := p.db.First(pembayaran, id).Error
	return pembayaran, err
}

func (p *pembayaranService) GetPembayaranByIdUser(id uint) (*[]domain.PembayaranUser, error) {
	pembayaran := &[]domain.PembayaranUser{}
	err := p.db.Where("id_meta_user = ?", id).Find(pembayaran).Error
	return pembayaran, err
}
