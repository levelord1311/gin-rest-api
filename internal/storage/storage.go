package storage

type Storage interface {
	FindById(id string) (*User, error)
}
