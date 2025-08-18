package handling

import (
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type HandlerEntry struct {
	Prefix  string
	Handler interface{}
}

type Route struct {
	httpMethod    string         // HTTP-метод (GET, POST и т.д.)
	prefix        string         // Префикс URL
	handlerName   string         // Имя обработчика
	actionName    string         // Имя действия
	expression    regexp.Regexp  // Регулярное выражение для маршрута
	handlerMethod reflect.Method // Метод обработчика через рефлексию
}

var httpMethods = []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut}

func generateRoutes(entries ...HandlerEntry) []Route {
	routes := make([]Route, 0, 10)
	for _, entry := range entries {
		handlerType := reflect.TypeOf(entry.Handler)
		innerMethods := getAnonymousFieldMethods(handlerType)
		for i := 0; i < handlerType.NumMethod(); i++ {
			handlerMethod := handlerType.Method(i)
			handlerMethodName := strings.ToUpper(handlerMethod.Name)
			for _, httpMethod := range httpMethods {
				if strings.HasPrefix(handlerMethodName, strings.ToUpper(httpMethod)) {
					if matchesPromotedMethodName(handlerMethod, innerMethods) {
						continue
					}
					route := Route{
						httpMethod:    httpMethod,
						prefix:        entry.Prefix,
						handlerName:   strings.Split(handlerType.Name(), "Handler")[0],
						actionName:    strings.Split(handlerMethodName, httpMethod)[1],
						handlerMethod: handlerMethod,
					}
					generateRegularExpression(entry.Prefix, &route)
					routes = append(routes, route)
				}
			}
		}
	}

	return routes
}

func matchesPromotedMethodName(targetMethod reflect.Method, httpMethods []reflect.Method) bool {
	for _, m := range httpMethods {
		if targetMethod.Name == m.Name {
			return true
		}
	}
	return false
}

func getAnonymousFieldMethods(target reflect.Type) []reflect.Method {
	methods := []reflect.Method{}
	for i := 0; i < target.NumField(); i++ {
		field := target.Field(i)
		if field.Anonymous && field.IsExported() {
			for j := 0; j < field.Type.NumMethod(); j++ {
				method := field.Type.Method(j)
				if method.IsExported() {
					methods = append(methods, method)
				}
			}
		}
	}
	return methods
}

func generateRegularExpression(prefix string, route *Route) {
	if prefix != "" && !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	pattern := "(?i)" + "/" + prefix + route.actionName
	if route.httpMethod == http.MethodGet {
		for i := 1; i < route.handlerMethod.Type.NumIn(); i++ {
			if route.handlerMethod.Type.In(i).Kind() == reflect.Int {
				pattern += "/([0-9]*)"
			} else {
				pattern += "/([A-z0-9]*)"
			}
		}
	}
	pattern = "^" + pattern + "[/]?$"
	route.expression = *regexp.MustCompile(pattern)
}
