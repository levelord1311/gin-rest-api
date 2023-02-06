package handlers

import (
	"gin-rest-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userPath = "/user"
)

type Handler struct {
	service Service
}

type Service interface {
	CreateUser(name string) (string, error)
	GetUser(id string) (*model.User, error)
}

func (h *Handler) Register(router *gin.Engine) {
	userEndpoints := router.Group(userPath)
	{
		userEndpoints.GET("/:id", h.GetUser)
		userEndpoints.POST("", h.CreateUser)
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Param("id")
	user, err := h.service.GetUser(id)
	status := http.StatusOK
	if err != nil {
		status = http.StatusNotFound
	}
	c.JSON(status, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	user := &model.User{}
	if err := c.BindJSON(user); err != nil {
		sendErrBadRequest(c, err)
		return
	}
	createdID, err := h.service.CreateUser(user.Name)
	if err != nil {
		sendErrBadRequest(c, err)
		return
	}
	c.JSON(http.StatusCreated, createdID)
}

func sendErrBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
