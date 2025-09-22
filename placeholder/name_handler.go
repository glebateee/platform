package placeholder

import (
	"fmt"
	"platform/http/actionresults"
	"platform/http/handling"
	"platform/logging"
	"platform/validation"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

type NameHandler struct {
	logging.Logger
	handling.URLGenerator
	validation.Validator
}

func (nh NameHandler) GetName(i int) actionresults.ActionResult {
	nh.Logger.Debugf("GetName method invoked with argument: %v", i)
	response := "Index out of bounds"
	if i < len(names) {
		response = fmt.Sprintf("Name #%v: %v", i, names[i])
	}
	return actionresults.NewTemplateAction("simple_message.html", response)
}

func (nh NameHandler) GetNames() actionresults.ActionResult {
	nh.Logger.Debug("GetNames method invoked")
	return actionresults.NewTemplateAction("simple_message.html", names)
}

type NewName struct {
	Name          string `validation:"required,min:3"`
	InsertAtStart bool
}

func (n NameHandler) GetForm() actionresults.ActionResult {
	postUrl, _ := n.URLGenerator.GenerateUrl(NameHandler.PostName)
	return actionresults.NewTemplateAction("name_form.html", postUrl)
}

func (nh NameHandler) PostName(new NewName) actionresults.ActionResult {
	nh.Logger.Debugf("PostName method invoked with argument %v", new)
	if ok, errs := nh.Validator.Validate(&new); !ok {
		return actionresults.NewTemplateAction("validation_errors.html", errs)
	}
	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}
	return nh.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) GetJsonData() actionresults.ActionResult {
	return actionresults.NewJsonAction(names)
}

func (n NameHandler) GetRedirect() actionresults.ActionResult {
	return n.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) redirectOrError(handler interface{}) actionresults.ActionResult {
	url, err := n.GenerateUrl(handler)
	if err == nil {
		return actionresults.NewRedirectAction(url)
	} else {
		return actionresults.NewErrorAction(err)
	}
}
