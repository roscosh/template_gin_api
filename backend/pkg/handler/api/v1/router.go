package v1

import (
	"github.com/gin-gonic/gin"
	"template_gin_api/pkg/handler/api/base_api"
	"template_gin_api/pkg/handler/api/v1/auth"
	"template_gin_api/pkg/handler/api/v1/users"
	"template_gin_api/pkg/service"
)

type V1Router struct {
	*baseApi.BaseAPIRouter
	service *service.Service
}

func NewV1Router(
	baseAPIHandler *baseApi.BaseAPIRouter,
	service *service.Service,
) *V1Router {
	return &V1Router{
		BaseAPIRouter: baseAPIHandler,
		service:       service,
	}
}

func (h *V1Router) RegisterRoutes(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	authRouter := auth.NewAuthRouter(h.BaseAPIRouter, h.service.Authorization, h.service.Users)
	authRouter.RegisterRoutes(authGroup)

	usersGroup := router.Group("/users", h.Middleware.AuthRequired)
	usersRouter := users.NewUsersRouter(h.BaseAPIRouter, h.service.Users)
	usersRouter.RegisterRoutes(usersGroup)
}
