package service

import (
	"order-go/models"
	"order-go/repository"
)

type UserService interface {
	GetByLoginName(loginName string) (models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func (u userService) GetByLoginName(loginName string) (models.User, error) {
	return u.repo.SelectOne(loginName)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
