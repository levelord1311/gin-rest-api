package service

import "gin-rest-api/internal/storage"

type Service interface {
	CreateUser(name string) (string, error)
	GetUser(id string) (*storage.User, error)
}

type userService struct {
	storage storage.Storage
}

func (u *userService) CreateUser(name string) (string, error) {
	createdID, err := u.storage.Create(name)
	if err != nil {
		return "", err
	}
	return createdID, nil
}

func (u *userService) GetUser(id string) (*storage.User, error) {
	user, err := u.storage.FindById(id)

	return user, err
}
