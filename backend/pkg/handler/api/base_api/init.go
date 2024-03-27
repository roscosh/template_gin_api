package baseApi

import (
	"github.com/gin-gonic/gin"
	"template_gin_api/misc"
	"template_gin_api/pkg/service"
)

var logger = misc.GetLogger()

type Router struct {
	Middleware *Middleware
	Config     *misc.Config
}

func NewRouter(services *service.Service, config *misc.Config) *Router {
	return &Router{
		Middleware: NewMiddleware(services.Middleware, config),
		Config:     config,
	}
}

type ApiRouter interface {
	RegisterHandlers(router *gin.RouterGroup)
}
