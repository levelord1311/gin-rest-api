package handlers

import (
	"gin-rest-api/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	userPath = "/user"
)

type Handler struct {
	service storage.Service
}

func (h *Handler) Register(router *gin.Engine) {
	userEndpoints := router.Group(userPath)
	{
		userEndpoints.GET("/:id", h.GetUser)
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
