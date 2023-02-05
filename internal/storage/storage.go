package storage

type Storage interface {
	Create(name string) (string, error)
	FindById(id string) (*User, error)
}
