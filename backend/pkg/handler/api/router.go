package api

import (
	"github.com/gin-gonic/gin"
	"template_gin_api/pkg/handler/api/base_api"
	"template_gin_api/pkg/handler/api/v1"
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
	v1Group := router.Group("/v1", h.Middleware.SessionRequired)
	v1Router := v1.NewRouter(h.Router, h.service)
	v1Router.RegisterHandlers(v1Group)
}
