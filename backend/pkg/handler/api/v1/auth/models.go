package auth

import (
	"template_gin_api/misc/session"
	"template_gin_api/pkg/repository/sql"
)

type loginForm struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type loginResponse struct {
	*session.Session
}

type logoutResponse struct {
	*session.Session
}

type meResponse struct {
	*session.Session
}

type signUpForm struct {
	*sql.CreateUser
}

type signUpResponse struct {
	*sql.User
}
