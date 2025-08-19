package actionresults

type ErrorActionResult struct {
	error
}

func NewErrorAction(err error) ActionResult {
	return &ErrorActionResult{error: err}
}

func (action *ErrorActionResult) Execute(ctx *ActionContext) error {
	return action.error
}
