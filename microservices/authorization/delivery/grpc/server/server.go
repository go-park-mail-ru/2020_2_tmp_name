package server

import (
	"context"
	auth "park_2020/2020_2_tmp_name/microservices/authorization/delivery/grpc/protobuf"
	"sync"
)

type SessionManager struct {
	mu          sync.RWMutex
	userUsecase UserUsecase
}

func NewSessionManager(u UserUsecase) *SessionManager {
	return &SessionManager{
		mu:          sync.RWMutex{},
		userUsecase: u,
	}
}

func (sm *SessionManager) Login(ctx context.Context, data *auth.LoginData) (*auth.Session, error) {
	var err error
	var session *auth.Session
	sm.mu.Lock()
	session.Sess, err = sm.userUsecase.Login(*data)
	sm.mu.Unlock()
	return session, err
}

func (sm *SessionManager) Logout(ctx context.Context, session *auth.Session) error {
	return sm.userUsecase.Logout(session.Sess)
}
