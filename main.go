package main

import (
	"platform/config"
	"platform/logging"
)

func writeMessage(logger logging.Logger, cfg config.Configuration) {
	if section, ok := cfg.GetSection("main"); ok {
		if message, ok := section.GetString("message"); ok {
			logger.Info(message)
		} else {
			logger.Panic("Cannot find configuration setting")
		}
	} else {
		logger.Panic("Config section not found")
	}
}
func main() {
	var cfg config.Configuration
	var err error
	if cfg, err = config.Load("config.json"); err != nil {
		panic(err)
	}
	var logger logging.Logger = logging.NewDefaultLogger(cfg)
	writeMessage(logger, cfg)
}
