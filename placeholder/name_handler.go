package placeholder

import (
	"fmt"
	"platform/logging"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

type NameHandler struct {
	logging.Logger
}

func (nh NameHandler) GetName(i int) string {
	nh.Logger.Debugf("GetName method invoked with argument: %v", i)
	if i < len(names) {
		return fmt.Sprintf("Name #%v: %v", i, names[i])
	} else {
		return "Index out of bounds"
	}
}

func (nh NameHandler) GetNames() string {
	nh.Logger.Debug("GetNames method invoked")
	return fmt.Sprintf("Names: %v", names)
}

type NewName struct {
	Name          string
	InsertAtStart bool
}

func (nh NameHandler) PostName(new NewName) string {
	nh.Logger.Debugf("PostName method invoked with argument %v", new)
	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}
	return fmt.Sprintf("Names: %v", names)
}
