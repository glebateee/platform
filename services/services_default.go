package services

import (
	"platform/config"
	"platform/logging"
)

func RegisterDefaultServices() {
	err := addSingleton(func() (c config.Configuration) {
		c, loadErr := config.Load("config.json")
		if loadErr != nil {
			panic(loadErr)
		}
		return
	})
	if err != nil {
		panic(err)
	}
	err = addSingleton(func(c config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(c)
	})
	if err != nil {
		panic(err)
	}
}
