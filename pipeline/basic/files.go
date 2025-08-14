package basic

import (
	"net/http"
	"platform/config"
	"platform/pipeline"
	"strings"
)

type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
	Config        config.Configuration
}

func (sfc *StaticFileComponent) Init() {
	// var cfg config.Configuration
	// err := services.GetService(&cfg)
	sfc.urlPrefix = sfc.Config.GetStringDefault("files:urlprefix", "/files/")
	if path, ok := sfc.Config.GetString("files:path"); ok {
		sfc.stdLibHandler = http.StripPrefix(sfc.urlPrefix, http.FileServer(http.Dir(path)))
	} else {
		panic("Cannot load file configuration settings")
	}

}

func (sfc *StaticFileComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	if !strings.EqualFold(ctx.Request.URL.Path, sfc.urlPrefix) &&
		strings.HasPrefix(ctx.Request.URL.Path, sfc.urlPrefix) {
		sfc.stdLibHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
