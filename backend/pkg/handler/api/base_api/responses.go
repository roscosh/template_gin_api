package baseApi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Error: err.Error()}
}

func Response401(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, NewErrorResponse(err))
}

func Response403(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, NewErrorResponse(err))
}

func Response404(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, NewErrorResponse(err))
}
