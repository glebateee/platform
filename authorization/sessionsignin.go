package authorization

import (
	"context"
	"platform/authorization/identity"
	"platform/services"
	"platform/sessions"
)

const USER_SESSION_KEY string = "USER"

type SessionSignInMgr struct {
	context.Context
}

func (s *SessionSignInMgr) getSession() (session sessions.Session, err error) {
	err = services.GetServiceForContext(s.Context, &session)
	return
}

// SignIn implements identity.SignInManager.
func (s *SessionSignInMgr) SignIn(user identity.User) error {
	session, err := s.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, user.GetID())
	}
	return err
}

// SignOut implements identity.SignInManager.
func (s *SessionSignInMgr) SignOut(user identity.User) error {
	session, err := s.getSession()
	if err == nil {
		session.SetValue(USER_SESSION_KEY, nil)
	}
	return err
}

func RegisterDefaultSignInService() {
	err := services.AddScoped(func(ctx context.Context) identity.SignInManager {
		return &SessionSignInMgr{Context: ctx}
	})
	if err != nil {
		panic(err)
	}
}
