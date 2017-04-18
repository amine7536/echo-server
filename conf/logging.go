package conf

import (
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/amine7536/echo-server/utils"
)

// LoggingConfig specifies all the parameters needed for logging
type LoggingConfig struct {
	Level  string
	File   string
	Format string
}

// ConfigureLogging will take the logging configuration and also adds
// a few default parameters
func ConfigureLogging(config *LoggingConfig) (*logrus.Entry, error) {

	// Get the host running the app
	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	// Use a log file if defined in the configuration file
	if config.File != "" {
		if utils.IsValidPath(config.File) {
			f, errOpen := os.OpenFile(config.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
			if errOpen != nil {
				return nil, errOpen
			}
			logrus.SetOutput(f)
		}
	}

	if config.Level != "" {
		level, err := logrus.ParseLevel(strings.ToUpper(config.Level))
		if err != nil {
			return nil, err
		}
		logrus.SetLevel(level)
	}

	switch config.Format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:    true,
			DisableTimestamp: false,
		})
	}

	return logrus.StandardLogger().WithField("host", host), nil
}
