package v1

import (
	"github.com/gin-gonic/gin"
	"template_gin_api/pkg/handler/api/base_api"
	"template_gin_api/pkg/handler/api/v1/auth"
	"template_gin_api/pkg/handler/api/v1/users"
	"template_gin_api/pkg/service"
)

type router struct {
	*baseApi.Router
	service *service.Service
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	service *service.Service,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	authRouter := auth.NewRouter(h.Router, h.service.Authorization)
	authRouter.RegisterHandlers(authGroup)

	usersGroup := router.Group("/users", h.Middleware.AdminRequired)
	usersRouter := users.NewRouter(h.Router, h.service.Users)
	usersRouter.RegisterHandlers(usersGroup)

}
