package domain

import (
	"context"
	"github.com/gin-gonic/gin"
)

type (
	UserService interface {
		GetOne(ctx context.Context, id int) (*User, error)
	}
	UserHandler interface {
		HandleUserGroup(group *gin.RouterGroup)
	}
)
