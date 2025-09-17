package templates

import "io"

type TemplateExecutor interface {
	ExecTemplate(writer io.Writer, name string, data interface{}) error
	ExecTemplateWithFunc(writer io.Writer, name string, data interface{}, handlerFunc InvokeHandlerFunc) error
}

type InvokeHandlerFunc func(handlerName string, methodName string, args ...interface{}) interface{}
