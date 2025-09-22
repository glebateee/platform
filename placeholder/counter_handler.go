package placeholder

import (
	"fmt"
	"platform/sessions"
)

type CounterHandler struct {
	sessions.Session
}

func (ch CounterHandler) GetCounter() string {
	counter := ch.Session.GetValueDefault("counter", 0).(int)
	ch.Session.SetValue("counter", counter+1)
	return fmt.Sprintf("Counter: %v", counter)
}
