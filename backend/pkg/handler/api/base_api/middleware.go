package baseApi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"template_gin_api/misc"
	"template_gin_api/misc/session"
	"template_gin_api/pkg/service"
)

const (
	userCtx = "userId"
)

type Middleware struct {
	service *service.MiddlewareService
	config  *misc.Config
}

func NewMiddleware(service *service.MiddlewareService, config *misc.Config) *Middleware {
	return &Middleware{service: service, config: config}
}

func (h *Middleware) SessionRequired(c *gin.Context) {
	token, _ := c.Cookie(session.CookieSessionName)
	sessionObj, err := h.service.GetExistSession(token)
	if err != nil {
		sessionObj, err = h.service.CreateSession()
		if err != nil {
			Response404(c, err)
			c.Abort()
			return
		}
	}
	c.Set(userCtx, sessionObj)

	SetCookie(c, sessionObj)
	c.Next()

	h.service.UpdateSession(sessionObj)
}

func (h *Middleware) AdminRequired(c *gin.Context) {
	sessionObj := GetSession(c)

	if !sessionObj.IsAdmin() {
		Response403(c, errors.New("Нужны права администратора для этого запроса!"))
		c.Abort()
		return
	}
}

func (h *Middleware) AuthRequired(c *gin.Context) {
	sessionObj := GetSession(c)

	if !sessionObj.IsAuthenticated() {
		Response401(c, errors.New("Нужно залогиниться для этого запроса!"))
		c.Abort()
		return
	}
}
