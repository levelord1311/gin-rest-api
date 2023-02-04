package handlers

import (
	"bytes"
	"encoding/json"
	"gin-rest-api/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockService struct {
	users map[string]storage.User
}

func (s *MockService) GetUser(id string) (*storage.User, error) {
	var err error
	user, ok := s.users[id]
	if !ok {
		err = storage.ErrNotFound
	}
	return &user, err
}

func (s *MockService) CreateUser(name string) (string, error) {
	if _, ok := s.users[name]; ok {
		return "", storage.ErrAlreadyExists
	}
	id := "2"
	s.users[id] = storage.User{
		ID:   id,
		Name: name,
	}
	return id, nil
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := map[string]storage.User{
		"1": {
			ID:   "1",
			Name: "Boris",
		},
		"2": {
			ID:   "2",
			Name: "Tommy",
		},
	}

	service := &MockService{users: store}
	h := &Handler{
		service,
	}

	cases := []struct {
		name   string
		sendID string
		want   storage.User
		code   int
	}{
		{
			name:   "Get existing user",
			sendID: "1",
			want: storage.User{
				ID:   "1",
				Name: "Boris",
			},
			code: http.StatusOK,
		},
		{
			name:   "Get another existing user",
			sendID: "2",
			want: storage.User{
				ID:   "2",
				Name: "Tommy",
			},
			code: http.StatusOK,
		},
		{
			name:   "Get user with non-existing ID",
			sendID: "3",
			want:   storage.User{},
			code:   http.StatusNotFound,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			request, _ := http.NewRequest(http.MethodGet, userPath, nil)
			c.Request = request

			c.Params = []gin.Param{
				{"id", test.sendID},
			}

			h.GetUser(c)

			assert.Equal(t, test.code, w.Code)

			got := storage.User{}
			json.Unmarshal(w.Body.Bytes(), &got)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := map[string]storage.User{
		"Turkish": storage.User{
			ID:   "1",
			Name: "Turkish",
		},
	}
	service := &MockService{users: store}
	h := Handler{service: service}

	cases := []struct {
		name     string
		postData storage.User
		code     int
	}{
		{
			"create new user",
			storage.User{
				Name: "Boris",
			},

			http.StatusCreated,
		},
		{
			"send empty user data",
			storage.User{},
			http.StatusBadRequest,
		},
		{
			"user already exists",
			storage.User{
				Name: "Turkish",
			},
			http.StatusBadRequest,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			postData, err := json.Marshal(test.postData)
			assert.NoError(t, err)

			request, _ := http.NewRequest(http.MethodPost, userPath, bytes.NewBuffer(postData))
			c.Request = request

			h.CreateUser(c)

			assert.Equal(t, test.code, w.Code)
		})
	}
}
