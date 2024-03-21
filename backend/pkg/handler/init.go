package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "template_gin_api/docs"
	"template_gin_api/misc"
	"template_gin_api/pkg/handler/api"
	"template_gin_api/pkg/handler/api/base_api"
	"template_gin_api/pkg/service"
)

func InitRoutes(service *service.Service, config *misc.Config) *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseAPIRouter := baseApi.NewBaseAPIRouter(service, config)

	apiGroup := router.Group("/api/")
	apiRouter := api.NewApiRouter(baseAPIRouter, service)
	apiRouter.RegisterRoutes(apiGroup)

	return router
}
