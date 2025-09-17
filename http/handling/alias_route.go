package handling

import (
	"fmt"
	"net/http"
	"platform/http/actionresults"
	"platform/services"
	"reflect"
	"regexp"
)

func (rc *RouterComponent) AddMethodAlias(srcUrl string, method interface{}, methodArgs ...interface{}) *RouterComponent {
	var urlGen URLGenerator
	services.GetService(&urlGen)
	url, err := urlGen.GenerateUrl(method, methodArgs...)
	if err == nil {
		return rc.AddUrlAlias(srcUrl, url)
	}
	panic(err)
}

func (rc *RouterComponent) AddUrlAlias(srcUrl, targetUrl string) *RouterComponent {
	aliasFunc := func(interface{}) actionresults.ActionResult {
		return actionresults.NewRedirectAction(targetUrl)
	}
	alias := Route{
		httpMethod:  http.MethodGet,
		handlerName: "Alias",
		actionName:  "Redirect",
		expression:  *regexp.MustCompile(fmt.Sprintf("^%v[/]?$", srcUrl)),
		handlerMethod: reflect.Method{
			Type: reflect.TypeOf(aliasFunc),
			Func: reflect.ValueOf(aliasFunc),
		},
	}
	rc.routes = append([]Route{alias}, rc.routes...)
	return rc
}
