package postgres

import (
	"github.com/adityarizkyramadhan/garbage-market/domain"
	"gorm.io/gorm"
)

type metaDataService struct {
	db *gorm.DB
}

func NewMetaDataService(db *gorm.DB) domain.ServiceMetaData {
	return &metaDataService{db}
}

//register
func (m *metaDataService) Register(user *domain.MetaUser) (*domain.MetaUser, error) {
	if err := m.db.Create(user).Error; err != nil {
		return &domain.MetaUser{}, err
	}
	return user, nil
}

//login
func (m *metaDataService) Login(email string) (*domain.MetaUser, error) {
	user := &domain.MetaUser{}
	err := m.db.Where("email = ? ", email).First(user).Error
	return user, err
}

func (m *metaDataService) GetUserById(id uint) (*domain.MetaUser, error) {
	user := &domain.MetaUser{}
	err := m.db.Preload("TokoSampah").Preload("BasketJualans").Preload("PembayaranUsers").First(user, id).Error
	return user, err
}
