package sessions

import (
	"context"
	"platform/config"
	"platform/pipeline"
	"time"

	gorilla "github.com/gorilla/sessions"
)

type SessionComponent struct {
	store *gorilla.CookieStore
	config.Configuration
}

func (sc *SessionComponent) Init() {
	cookieKey, found := sc.Configuration.GetString("sessions:key")
	if !found {
		panic("Session key not found in configuration")
	}
	if sc.Configuration.GetBoolDefault("sessions:cyclekey", true) {
		cookieKey += time.Now().String()
	}
	sc.store = gorilla.NewCookieStore([]byte(cookieKey))
}

func (sc *SessionComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	session, _ := sc.store.Get(ctx.Request, SESSION__CONTEXT_KEY)
	ctxWithSession := context.WithValue(ctx.Request.Context(), SESSION__CONTEXT_KEY, session)
	ctx.Request = ctx.Request.WithContext(ctxWithSession)
	next(ctx)
	sc.store.Save(ctx.Request, ctx.ResponseWriter, session)
}
