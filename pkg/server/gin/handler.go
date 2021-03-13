package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/niqinge/test-simple/pkg/service"
)

type Handler struct {
	Manager *service.Manager
}

func (h *Handler) AddUser () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *Handler) DelUser () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *Handler) UpdateUser () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *Handler) QueryUser () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *Handler) QueryUserList () gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}