package service

import (
	"gin-rest-api/internal/model"
)

type Storage interface {
	Create(name string) (string, error)
	FindById(id string) (*model.User, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (u *Service) CreateUser(name string) (string, error) {
	createdID, err := u.storage.Create(name)
	if err != nil {
		return "", err
	}
	return createdID, nil
}

func (u *Service) GetUser(id string) (*model.User, error) {
	user, err := u.storage.FindById(id)

	return user, err
}
