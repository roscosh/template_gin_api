package baseApi

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"template_gin_api/misc/session"
)

func GetPathID(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.New("id должен быть числом!")
	}
	return id, err
}

func SetCookie(c *gin.Context, sessionObj *session.Session) {
	c.SetCookie(session.CookieSessionName, sessionObj.Token, sessionObj.Expires, "", "", false, false)
}
