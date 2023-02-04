package service

import "gin-rest-api/internal/storage"

type Service interface {
	CreateUser(name string) (string, error)
	GetUser(id string) (*storage.User, error)
}
