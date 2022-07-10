package postgres

import (
	"github.com/adityarizkyramadhan/garbage-market/domain"
	"gorm.io/gorm"
)

type basketService struct {
	db *gorm.DB
}

//NewbasketService
func NewbasketService(db *gorm.DB) domain.BasketJualanService {
	return &basketService{db}
}

func (b *basketService) CreateBasketJualan(basket *domain.BasketJualan) error {
	return b.db.Create(basket).Error
}

func (b *basketService) GetBasketJualanById(id uint) (*domain.BasketJualan, error) {
	basket := &domain.BasketJualan{}
	err := b.db.First(basket, id).Error
	return basket, err
}

//delete basket jualan
func (b *basketService) DeleteBasketJualan(idBusket uint) error {
	basket := &domain.BasketJualan{}
	err := b.db.Where("id_busket = ?", idBusket).Delete(basket).Error
	return err
}

func (b *basketService) GetBasketJualanByIdUser(idUser uint) (*[]domain.BasketJualan, error) {
	basket := &[]domain.BasketJualan{}
	err := b.db.Where("id_user = ?", idUser).Find(basket).Error
	return basket, err
}

func (b *basketService) UpdateBasketJualan(id uint, basket *domain.BasketJualan) error {
	return b.db.Model(&domain.BasketJualan{}).Where("id = ?", id).Save(basket).Error
}
