package mysqlDB

import (
	"database/sql"
	"gin-rest-api/internal/storage"
	_ "github.com/go-sql-driver/mysql"
)

var _ storage.Storage = &db{}

type db struct {
	db *sql.DB
}

func NewStorage(storage *sql.DB) *db {
	return &db{
		db: storage,
	}
}

func (d *db) Create(name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) FindById(id string) (*storage.User, error) {
	//TODO implement me
	panic("implement me")
}
