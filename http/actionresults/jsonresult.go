package actionresults

import "encoding/json"

type JsonActionResult struct {
	data interface{}
}

func NewJsonAction(data interface{}) ActionResult {
	return &JsonActionResult{data: data}
}

func (action *JsonActionResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.ResponseWriter)
	return encoder.Encode(action.data)
}
