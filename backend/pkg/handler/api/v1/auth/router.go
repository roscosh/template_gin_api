package auth

import (
	"github.com/gin-gonic/gin"
	"template_gin_api/pkg/handler/api/base_api"
	"template_gin_api/pkg/service"
)

type router struct {
	*baseApi.Router
	service *service.AuthService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	service *service.AuthService,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: service,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/me", h.me)
	router.POST("/sign_up", h.signUp)
	router.POST("/login", h.login)
	router.POST("/logout", h.Middleware.AuthRequired, h.logout)
}
