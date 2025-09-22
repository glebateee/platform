package authorization

import (
	"net/http"
	"platform/authorization/identity"
	"platform/config"
	"platform/http/handling"
	"platform/pipeline"
	"regexp"
	"strings"
)

type AuthMiddlewareComponent struct {
	prefix    string
	condition identity.AuthorizationCondition
	pipeline.RequestPipeline
	config.Configuration
	authFailUrl string
	fallbacks   map[*regexp.Regexp]string
}

func NewAuthComponent(prefix string, condition identity.AuthorizationCondition, requestHandlers ...interface{}) *AuthMiddlewareComponent {
	entries := []handling.HandlerEntry{}
	for _, handler := range requestHandlers {
		entries = append(entries, handling.HandlerEntry{Prefix: prefix, Handler: handler})
	}
	router := handling.NewRouter(entries...)
	return &AuthMiddlewareComponent{
		prefix:          "/" + prefix,
		condition:       condition,
		RequestPipeline: pipeline.CreatePipeline(router),
		fallbacks:       map[*regexp.Regexp]string{},
	}
}

func (*AuthMiddlewareComponent) Init()                                 {}
func (*AuthMiddlewareComponent) ImplementsProcessRequestWithServices() {}

func (cmp *AuthMiddlewareComponent) ProcessRequestWithServices(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext),
	user identity.User) {
	if strings.HasPrefix(ctx.Request.URL.Path, cmp.prefix) {
		for expr, target := range cmp.fallbacks {
			if expr.MatchString(ctx.Request.URL.Path) {
				http.Redirect(ctx.ResponseWriter, ctx.Request, target, http.StatusSeeOther)
				return
			}
		}
		if cmp.condition.Validate(user) {
			cmp.RequestPipeline.ProcessRequest(ctx.Request, ctx.ResponseWriter)
		} else {
			if cmp.authFailUrl != "" {
				http.Redirect(ctx.ResponseWriter, ctx.Request, cmp.authFailUrl, http.StatusSeeOther)
			} else if user.IsAuthenticated() {
				ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
			} else {
				ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)

			}
		}
	} else {
		next(ctx)
	}
}
