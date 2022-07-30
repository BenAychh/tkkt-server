package usrhndlr

import (
	"github.com/gin-gonic/gin"
	"tkkt-auth/src/domain"
)

type Handler struct {
	UserService domain.UserService
}

func New(userService domain.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}

func (h *Handler) HandleUserGroup(group *gin.RouterGroup) {

}
