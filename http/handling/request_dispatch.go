package handling

import (
	"fmt"
	"io"
	"net/http"
	"platform/http/handling/params"
	"platform/pipeline"
	"platform/services"
	"reflect"
	"strings"
)

type RouterComponent struct {
	routes []Route
}

func NewRouter(handlers ...HandlerEntry) *RouterComponent {
	return &RouterComponent{generateRoutes(handlers...)}
}

func (router *RouterComponent) Init() {}

func (router *RouterComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	for _, route := range router.routes {
		if strings.EqualFold(ctx.Request.Method, route.httpMethod) {
			matches := route.expression.FindAllStringSubmatch(ctx.URL.Path, -1)
			if len(matches) > 0 {
				rawParamVals := []string{}
				if len(matches[0]) > 1 {
					rawParamVals = matches[0][1:]
				}
				err := router.invokeHandler(route, rawParamVals, ctx)
				if err == nil {
					next(ctx)
				} else {
					ctx.Error(err)
				}
				return
			}
		}
	}
	ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
}

func printRoute(route Route) {
	fmt.Printf("\n-------------\nhttpMethod: %v\n", route.httpMethod)
	fmt.Printf("prefix: %v\n", route.prefix)
	fmt.Printf("handlerName: %v\n", route.handlerName)
	fmt.Printf("actionName: %v\n", route.actionName)
	fmt.Printf("actionName: %v\n", route.expression.String())
	fmt.Printf("handlerMethod: %v\n-------------\n", route.handlerMethod.Name)

}

func (router *RouterComponent) invokeHandler(route Route, rawParams []string, ctx *pipeline.ComponentContext) error {
	paramVals, err := params.GetParametersFromRequest(ctx.Request, route.handlerMethod, rawParams)
	if err == nil {
		structVal := reflect.New(route.handlerMethod.Type.In(0))
		services.PopulateForContext(ctx.Context(), structVal.Interface())
		paramVals = append([]reflect.Value{structVal.Elem()}, paramVals...)
		result := route.handlerMethod.Func.Call(paramVals)
		io.WriteString(ctx.ResponseWriter, fmt.Sprint(result[0].Interface()))
		printRoute(route)
	}
	return err
}
