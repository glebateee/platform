package pipeline

import (
	"net/http"
	"platform/services"
	"reflect"
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

func CreatePipeline(components ...interface{}) RequestPipeline {
	f := emptyPipeline
	for i := len(components) - 1; i >= 0; i-- {
		nf := f
		currentComponent := components[i]
		services.Populate(currentComponent)
		if servComp, ok := currentComponent.(ServicesMiddlwareComponent); ok {
			f = createServiceDependentFunction(currentComponent, nf)
			servComp.Init()
		} else if stdComp, ok := currentComponent.(MiddlewareComponent); ok {
			f = func(ctx *ComponentContext) {
				if ctx.error == nil {
					stdComp.ProcessRequest(ctx, nf)
				}
			}
			stdComp.Init()
		}

	}
	return f
}

func createServiceDependentFunction(component interface{}, nextFunc RequestPipeline) RequestPipeline {
	method := reflect.ValueOf(component).MethodByName("ProcessRequestWithServices")
	if method.IsValid() {
		return func(ctx *ComponentContext) {
			if ctx.error == nil {
				_, err := services.CallForContext(
					ctx.Request.Context(),
					method.Interface(),
					ctx,
					nextFunc,
				)
				if err != nil {
					ctx.Error(err)
				}
			}
		}
	} else {
		panic("No ProcessRequestWithServices method defined")
	}
}
