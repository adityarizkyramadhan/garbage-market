package postgres

import (
	"errors"

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
	data, err := t.GetTokoByIdUser(uint(toko.IdMetaUser))
	if err == nil {
		return err
	}
	if data.IdMetaUser != 0 {
		return errors.New("toko sampah sudah ada")
	}
	return t.db.Create(toko).Error
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
	err := t.db.Where("id_meta_user = ?", idUser).First(toko).Error
	return toko, err
}
