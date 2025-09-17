package actionresults

import (
	"platform/templates"
)

type TemplateActionResult struct {
	templateName string
	data         interface{}
	templates.TemplateExecutor
	templates.InvokeHandlerFunc
}

func NewTemplateAction(name string, data interface{}) ActionResult {
	return &TemplateActionResult{templateName: name, data: data}
}

func (action *TemplateActionResult) Execute(ctx *ActionContext) error {
	return action.TemplateExecutor.ExecTemplateWithFunc(ctx.ResponseWriter, action.templateName, action.data, action.InvokeHandlerFunc)
}
