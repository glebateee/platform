package placeholder

import (
	"fmt"
	"platform/http/actionresults"
	"platform/logging"
	"time"
)

type DayHandler struct {
	logging.Logger
}

// func (dh DayHandler) GetDay() string {
// 	return fmt.Sprintf("Day: %v", time.Now().Day())
// }

func (h DayHandler) GetDay() actionresults.ActionResult {
	return actionresults.NewTemplateAction("day_template.html", fmt.Sprintf("Day: %v", time.Now().Day()))
}
