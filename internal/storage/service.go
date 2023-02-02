package storage

type Service interface {
	GetUser(id string) (*User, error)
}
