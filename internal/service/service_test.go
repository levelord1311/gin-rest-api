package service

import (
	"gin-rest-api/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ storage.Storage = &stubStorage{}

type stubStorage map[string]*storage.User

func (s stubStorage) Create(name string) (string, error) {
	if _, ok := s[name]; ok {
		return "", storage.ErrAlreadyExists
	}
	return "2", nil
}

func (s stubStorage) FindById(id string) (*storage.User, error) {
	user, ok := s[id]
	if !ok {
		return nil, storage.ErrNotFound
	}
	return user, nil
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := stubStorage{
		"1": {
			"1",
			"Boris",
		},
		"2": {
			"2",
			"Turkish",
		},
	}

	service := &userService{
		storage: store,
	}

	cases := []struct {
		name   string
		sendID string
		want   *storage.User
		err    error
	}{
		{
			name:   "Get existing user",
			sendID: "1",
			want: &storage.User{
				ID:   "1",
				Name: "Boris",
			},
		},
		{
			name:   "Get another existing user",
			sendID: "2",
			want: &storage.User{
				ID:   "2",
				Name: "Turkish",
			},
		},
		{
			name:   "Get non-existing user",
			sendID: "3",
			want:   nil,
			err:    storage.ErrNotFound,
		},
	}
	for _, test := range cases {

		t.Run(test.name, func(t *testing.T) {
			got, err := service.GetUser(test.sendID)
			switch test.err {
			case storage.ErrNotFound:
				assert.EqualError(t, err, storage.ErrNotFound.Error())
			default:
				assert.NoError(t, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := stubStorage{
		"Boris": {
			"1",
			"Boris",
		},
	}

	service := &userService{
		storage: store,
	}

	cases := []struct {
		name     string
		userName string
		want     string
		err      error
	}{
		{
			name:     "Create new user",
			userName: "Turkish",
			want:     "2",
			err:      nil,
		},
		{
			name:     "User already exists",
			userName: "Boris",
			want:     "",
			err:      storage.ErrAlreadyExists,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got, err := service.CreateUser(test.userName)
			switch test.err {
			case storage.ErrAlreadyExists:
				assert.EqualError(t, err, storage.ErrAlreadyExists.Error())
			default:
				assert.NoError(t, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}
