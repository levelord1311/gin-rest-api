package storage

type Storage interface {
	Create(name string) (*User, error)
	FindById(id string) (*User, error)
}
