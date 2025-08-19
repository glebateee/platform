package placeholder

import (
	"fmt"
	"platform/http/actionresults"
	"platform/logging"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

type NameHandler struct {
	logging.Logger
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
	Name          string
	InsertAtStart bool
}

func (nh NameHandler) PostName(new NewName) actionresults.ActionResult {
	nh.Logger.Debugf("PostName method invoked with argument %v", new)
	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}
	return actionresults.NewRedirectAction("/names")
}

func (n NameHandler) GetJsonData() actionresults.ActionResult {
	return actionresults.NewJsonAction(names)
}
