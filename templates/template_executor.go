package templates

import "io"

type TemplateExecutor interface {
	ExecTemplate(io.Writer, string, interface{}) error
}
