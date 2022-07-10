package postgres

import (
	"github.com/adityarizkyramadhan/garbage-market/domain"
	"gorm.io/gorm"
)

type tokoSampahService struct {
	db *gorm.DB
}

func NewTokoSampahService(db *gorm.DB) *tokoSampahService {
	return &tokoSampahService{db}
}

//Create new toko sampah
func (t *tokoSampahService) CreateTokoSampah(toko *domain.TokoSampah) error {
	if t.db.Where("id_user = ?", toko.IdUser).First(toko).RowsAffected == 0 {
		return t.db.Create(toko).Error
	}
	return nil
}

func (t *tokoSampahService) GetTokoSampahById(id uint) (*domain.TokoSampah, error) {
	toko := &domain.TokoSampah{}
	err := t.db.First(toko, id).Error
	return toko, err
}

func (t *tokoSampahService) UpdateTokoSampah(toko *domain.TokoSampah) error {
	return t.db.Save(toko).Error
}

func (t *tokoSampahService) GetTokoByIdUser(idUser uint) (*domain.TokoSampah, error) {
	toko := &domain.TokoSampah{}
	err := t.db.Where("id_user = ?", idUser).First(toko).Error
	return toko, err
}
