package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	. "template_gin_api/misc/session"
	"template_gin_api/pkg/repository/sql"
)

const (
	salt = "cups_managment_relabs"
)

type AuthService struct {
	sql *sql.UsersSQL
}

func newAuthService(sql *sql.UsersSQL) *AuthService {
	return &AuthService{sql: sql}
}

func (s *AuthService) Login(login string, password string, session *Session) (*Session, error) {
	token := generatePasswordHash(password)
	user, err := s.sql.GetUserByCredentials(login, token)
	if err != nil {
		return nil, errors.New("Нет пользователя с таки логином/паролем!")
	}
	session.SetSession(user)
	return session, nil
}

func (s *AuthService) Logout(session *Session) *Session {
	session.ResetSession()
	return session
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
