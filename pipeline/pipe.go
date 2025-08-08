package pipeline

import (
	"net/http"
)

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(*ComponentContext) { /* do nothing */ }

func (pl RequestPipeline) ProcessRequest(req *http.Request, resp http.ResponseWriter) error {
	ctx := ComponentContext{
		Request:        req,
		ResponseWriter: resp,
	}
	pl(&ctx)
	return ctx.error
}

func CreatePipeline(components ...MiddlewareComponent) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		fn := f
		currentComponent := components[i]
		f = func(ctx *ComponentContext) {
			if ctx.error == nil {
				currentComponent.ProcessRequest(ctx, fn)
			}
		}
		currentComponent.Init()
	}
	return f
}
