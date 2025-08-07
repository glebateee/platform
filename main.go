package main

import (
	"platform/config"
	"platform/logging"
	"platform/services"
)

func writeMessage(logger logging.Logger, cfg config.Confuguration) {
	if section, ok := cfg.GetSection("main"); ok {
		if msg, ok := section.GetString("message"); ok {
			logger.Info(msg)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}
func main() {
	services.RegisterDefaultServices()
	// var cfg config.Confuguration
	// services.GetService(&cfg)
	// var logger logging.Logger
	// services.GetService(&logger)
	services.Call(writeMessage)
	val := struct {
		message string
		logging.Logger
	}{message: "Hello from the struct"}
	services.Populate(&val)
	val.Logger.Debug(val.message)
}
