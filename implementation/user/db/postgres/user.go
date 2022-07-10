package postgres

import (
	"github.com/adityarizkyramadhan/garbage-market/domain"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) domain.ServiceUser {
	return &userService{db}
}

//create new user
func (u *userService) CreateUser(user *domain.User) error {
	return u.db.Create(user).Error
}

//get user by id
func (u *userService) GetUserById(id uint) (*domain.User, error) {
	user := &domain.User{}
	err := u.db.Preload("TokoSampah").Preload("PembayaranUsers").Preload("BasketJualans").First(user, id).Error
	return user, err
}

//update user by id
func (u *userService) UpdateUser(id uint, user *domain.User) error {
	return u.db.Model(user).Where("id = ?", id).Save(user).Error
}
