package repository

import (
	"gorm.io/gorm"
	"order-go/models"
)

type UserRepository interface {
	SelectOne(loginName string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) SelectOne(loginName string) (models.User, error) {
	var userInfo models.User
	result := u.db.Model(&models.User{}).First(&userInfo).Where("login_name=?", loginName).Scan(&userInfo)
	err := result.Error
	return userInfo, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
