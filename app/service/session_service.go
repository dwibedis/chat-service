package service

import "github.com/dwibedis/chat-service/app/repository"

type Session struct {
	sessionRepo *repository.Session
}

func NewSession(sessionRepo *repository.Session) *Session {
	return &Session{sessionRepo:sessionRepo}
}

func CreateSession()  {
	
}
