package services

import "github.com/gorilla/sessions"

var Session SessionService

type SessionService struct {
	SessionStore sessions.Store
	SessionName  string
}

func NewSession(name string, store sessions.Store) SessionService {
	Session = SessionService{
		SessionStore: store,
		SessionName:  name,
	}

	return Session
}

func (s *SessionService) GetAuthorSessionKey() int8 {
	return 1
}

func (s *SessionService) GetRequestIdKey() int8 {
	return 2
}
