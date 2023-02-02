package handlers

import (
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
			code: 200,
		},
		{
			name:   "Get another existing user",
			sendID: "2",
			want: storage.User{
				ID:   "2",
				Name: "Tommy",
			},
			code: 200,
		},
		{
			name:   "Get user with non-existing ID",
			sendID: "3",
			want:   storage.User{},
			code:   404,
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			request, _ := http.NewRequest(http.MethodGet, "user", nil)
			c.Request = request

			c.Params = []gin.Param{
				{"id", test.sendID},
			}

			h.GetUser(c)

			got := storage.User{}
			json.Unmarshal(w.Body.Bytes(), &got)

			assert.Equal(t, test.code, w.Code)
			assert.Equal(t, test.want, got)
		})
	}
}
