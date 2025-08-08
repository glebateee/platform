package placeholder

import (
	"errors"
	"io"
	"platform/config"
	"platform/pipeline"
	"platform/services"
)

type SimpleMessageComponent struct{}

func (c *SimpleMessageComponent) Init() {}

func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	var cfg config.Configuration
	if err := services.GetService(&cfg); err != nil {
		ctx.Error(err)
		return
	}
	if msg, ok := cfg.GetString("main:message"); ok {
		io.WriteString(ctx.ResponseWriter, msg)
	} else {
		ctx.Error(errors.New("cannot find config setting"))
	}
	next(ctx)
}
