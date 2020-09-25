package log

import "github.com/sirupsen/logrus"

func Initialize(c Config) *logrus.Logger {
	logger := logrus.New()
	if len(c.Level) > 0 {
		if level, err := logrus.ParseLevel(c.Level); err == nil {
			logger.SetLevel(level)
		} else {
			logrus.Errorf("Can't parse LOG_LEVEL: %s.", c.Level)
		}
	}
	// MiddleWareLog as JSON instead of the default ASCII formatter.
	// logger.SetFormatter(&logrus.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	// logger.SetOutput(os.Stdout)

	// kibana: time:@timestamp msg:message
	formatter := logrus.JSONFormatter{
		// disable, as we set our own
		TimestampFormat: "2006-01-02T15:04:05.000Z0700",
		/*
			DisableTimestamp: false,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "@level",
				logrus.FieldKeyMsg:   "@message",
				logrus.FieldKeyFunc:  "@caller",
			}, */
	}
	if len(c.TimestampFormat) > 0 {
		formatter.TimestampFormat = c.TimestampFormat
	}
	if c.FieldMap != nil {
		formatter.FieldMap = *c.FieldMap
	}
	logger.SetFormatter(&formatter)
	return logger
}
