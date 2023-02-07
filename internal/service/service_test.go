package service

import (
	"fmt"
	"gin-rest-api/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ Storage = &stubStorage{}

type stubStorage map[string]*model.User

func (s stubStorage) Create(name string) (string, error) {
	if _, ok := s[name]; ok {
		return "", fmt.Errorf("already exists")
	}
	return "2", nil
}

func (s stubStorage) FindById(id string) (*model.User, error) {
	user, ok := s[id]
	if !ok {
		return nil, fmt.Errorf("not found")
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

	service := NewService(store)

	cases := []struct {
		name    string
		sendID  string
		want    *model.User
		mustErr bool
	}{
		{
			name:   "Get existing user",
			sendID: "1",
			want: &model.User{
				ID:   "1",
				Name: "Boris",
			},
		},
		{
			name:   "Get another existing user",
			sendID: "2",
			want: &model.User{
				ID:   "2",
				Name: "Turkish",
			},
		},
		{
			name:    "Get non-existing user",
			sendID:  "3",
			want:    nil,
			mustErr: true,
		},
	}
	for _, test := range cases {

		t.Run(test.name, func(t *testing.T) {
			got, err := service.GetUser(test.sendID)
			if test.mustErr {
				assert.Error(t, err)
			} else {
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

	service := NewService(store)

	cases := []struct {
		name     string
		userName string
		want     string
		mustErr  bool
	}{
		{
			name:     "Create new user",
			userName: "Turkish",
			want:     "2",
		},
		{
			name:     "User already exists",
			userName: "Boris",
			want:     "",
			mustErr:  true,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got, err := service.CreateUser(test.userName)
			if test.mustErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}
