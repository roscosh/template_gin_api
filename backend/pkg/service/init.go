package service

import (
	"template_gin_api/misc"
	"template_gin_api/pkg/repository/redis"
	"template_gin_api/pkg/repository/sql"
)

var logger = misc.GetLogger()

type Service struct {
	Middleware    *MiddlewareService
	Authorization *AuthService
	Users         *UsersService
}

func NewService(sql *sql.SQL, redis *redis.Redis) *Service {
	return &Service{
		Middleware:    newMiddlewareService(sql.Users, redis.Session),
		Authorization: newAuthService(sql.Users),
		Users:         newUsersService(sql.Users),
	}
}
