package service

import (
	"gin-rest-api/internal/model"
)

type Storage interface {
	Create(name string) (string, error)
	FindById(id string) (*model.User, error)
}

type userService struct {
	storage Storage
}

func (u *userService) CreateUser(name string) (string, error) {
	createdID, err := u.storage.Create(name)
	if err != nil {
		return "", err
	}
	return createdID, nil
}

func (u *userService) GetUser(id string) (*model.User, error) {
	user, err := u.storage.FindById(id)

	return user, err
}
