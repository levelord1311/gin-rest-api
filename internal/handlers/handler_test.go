package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-rest-api/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var _ Service = &mockService{}

type mockService struct {
	users map[string]model.User
}

func (s *mockService) GetUser(id string) (*model.User, error) {
	var err error
	user, ok := s.users[id]
	if !ok {
		err = fmt.Errorf("not found")
	}
	return &user, err
}

func (s *mockService) CreateUser(name string) (string, error) {
	if _, ok := s.users[name]; ok {
		return "", fmt.Errorf("already exists")
	}
	id := "2"
	s.users[id] = model.User{
		ID:   id,
		Name: name,
	}
	return id, nil
}

func TestGetUser(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	store := map[string]model.User{
		"1": {
			ID:   "1",
			Name: "Boris",
		},
		"2": {
			ID:   "2",
			Name: "Tommy",
		},
	}

	service := &mockService{users: store}
	h := &Handler{
		service,
	}

	cases := []struct {
		name   string
		sendID string
		want   model.User
		code   int
	}{
		{
			name:   "Get existing user",
			sendID: "1",
			want: model.User{
				ID:   "1",
				Name: "Boris",
			},
			code: http.StatusOK,
		},
		{
			name:   "Get another existing user",
			sendID: "2",
			want: model.User{
				ID:   "2",
				Name: "Tommy",
			},
			code: http.StatusOK,
		},
		{
			name:   "Get user with non-existing ID",
			sendID: "3",
			want:   model.User{},
			code:   http.StatusNotFound,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			request, _ := http.NewRequest(http.MethodGet, userPath, nil)
			c.Request = request

			c.Params = []gin.Param{
				{"id", test.sendID},
			}

			h.GetUser(c)

			assert.Equal(t, test.code, w.Code)

			got := model.User{}
			json.Unmarshal(w.Body.Bytes(), &got)

			assert.Equal(t, test.want, got)
		})
	}
}

func TestCreateUser(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)

	store := map[string]model.User{
		"Turkish": {
			ID:   "1",
			Name: "Turkish",
		},
	}
	service := &mockService{users: store}
	h := Handler{service: service}

	cases := []struct {
		name     string
		postData model.User
		code     int
	}{
		{
			"create new user",
			model.User{
				Name: "Boris",
			},

			http.StatusCreated,
		},
		{
			"create another user",
			model.User{
				Name: "Tommy",
			},

			http.StatusCreated,
		},
		{
			"send empty user data",
			model.User{},
			http.StatusBadRequest,
		},
		{
			"user already exists",
			model.User{
				Name: "Turkish",
			},
			http.StatusBadRequest,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

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
