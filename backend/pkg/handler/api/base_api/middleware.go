package baseApi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"template_gin_api/misc"
	"template_gin_api/misc/session"
	"template_gin_api/pkg/service"
)

const (
	UserCtx = "userId"
)

type Handler struct {
	service *service.MiddlewareService
	config  *misc.Config
}

func NewHandler(service *service.MiddlewareService, config *misc.Config) *Handler {
	return &Handler{service: service, config: config}
}

func (h *Handler) GetSession(c *gin.Context) {
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
	c.Set(UserCtx, sessionObj)

	SetCookie(c, sessionObj)
	c.Next()

	h.service.UpdateSession(sessionObj)
}

func (h *Handler) AdminRequired(c *gin.Context) {
	sessionObj := c.MustGet(UserCtx).(*session.Session)

	if !sessionObj.IsAdmin() {
		Response403(c, errors.New("Нужны права администратора для этого запроса!"))
		c.Abort()
		return
	}
}

func (h *Handler) AuthRequired(c *gin.Context) {
	sessionObj := c.MustGet(UserCtx).(*session.Session)

	if !sessionObj.IsAuthenticated() {
		Response401(c, errors.New("Нужно залогиниться для этого запроса!"))
		c.Abort()
		return
	}
}
