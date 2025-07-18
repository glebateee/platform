package logging

import (
	"log"
	"os"
	"platform/config"
	"strings"
)

func LogLevelFromString(str string) (level LogLevel) {
	switch strings.ToLower(str) {
	case "debug":
		level = Debug
	case "information":
		level = Information
	case "warning":
		level = Warning
	case "fatal":
		level = Fatal
	case "none":
		level = None
	default:
		level = Debug
	}
	return
}

func NewDefaultLogger(cfg config.Configuration) Logger {
	var level LogLevel = Debug
	if configLevelString, ok := cfg.GetString("logging:level"); ok {
		level = LogLevelFromString(configLevelString)
	}
	flags := log.Lmsgprefix | log.Ltime
	return &DefaultLogger{
		minLevel: level,
		loggers: map[LogLevel]*log.Logger{
			Trace:       log.New(os.Stdout, "TRACE ", flags),
			Debug:       log.New(os.Stdout, "DEBUG ", flags),
			Information: log.New(os.Stdout, "INFO ", flags),
			Warning:     log.New(os.Stdout, "WARN ", flags),
			Fatal:       log.New(os.Stdout, "FATAL ", flags),
		},
		triggerPanic: true,
	}
}
