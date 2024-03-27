package users

import (
	"github.com/gin-gonic/gin"
	"template_gin_api/pkg/handler/api/base_api"
	"template_gin_api/pkg/service"
)

type router struct {
	*baseApi.Router
	service *service.UsersService
}

func NewRouter(
	baseAPIRouter *baseApi.Router,
	usersService *service.UsersService,
) baseApi.ApiRouter {
	return &router{
		Router:  baseAPIRouter,
		service: usersService,
	}
}

func (h *router) RegisterHandlers(router *gin.RouterGroup) {
	router.GET("/get_all", h.getAll)
	router.DELETE("/:id", h.delete)
	router.PUT("/:id", h.edit)
	router.PUT("/change_password/:id", h.changePassword)
}
