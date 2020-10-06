package log

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var fieldConfig FieldConfig

func Initialize(c Config) *logrus.Logger {
	fieldConfig.FieldMap = c.FieldMap
	if len(c.Duration) > 0 {
		fieldConfig.Duration = c.Duration
	} else {
		fieldConfig.Duration = "duration"
	}
	if len(c.Fields) > 0 {
		fields := strings.Split(c.Fields, ",")
		fieldConfig.Fields = &fields
	}
	logger := logrus.New()

	// kibana: time:@timestamp msg:message
	formatter := logrus.JSONFormatter{TimestampFormat: "2006-01-02T15:04:05.000Z0700"}
	if len(c.TimestampFormat) > 0 {
		formatter.TimestampFormat = c.TimestampFormat
	}
	if c.Map != nil {
		formatter.FieldMap = *c.Map
	}
	x := &formatter
	logger.SetFormatter(x)
	logrus.SetFormatter(x)
	if len(c.Level) > 0 {
		if level, err := logrus.ParseLevel(c.Level); err == nil {
			logger.SetLevel(level)
			logrus.SetLevel(level)
		} else {
			logrus.Errorf("Can't parse LOG_LEVEL: %s.", c.Level)
		}
	}
	return logger
}
func DebugDuration(ctx context.Context, start time.Time, args ...interface{}) {
	LogDuration(ctx, logrus.DebugLevel, start, args)
}
func InfoDuration(ctx context.Context, start time.Time, args ...interface{}) {
	LogDuration(ctx, logrus.InfoLevel, start, args)
}
func LogDuration(ctx context.Context, level logrus.Level, start time.Time, args ...interface{}) {
	if logrus.IsLevelEnabled(level) {
		end := time.Now()
		duration := end.Sub(start)
		fields := AppendFields(ctx, logrus.Fields{})
		fields[fieldConfig.Duration] = duration.Milliseconds()
		logrus.WithFields(fields).Log(level, args...)
	}
}
func LogfDuration(ctx context.Context, level logrus.Level, start time.Time, format string, args ...interface{}) {
	if logrus.IsLevelEnabled(level) {
		end := time.Now()
		duration := end.Sub(start)
		fields := AppendFields(ctx, logrus.Fields{})
		fields[fieldConfig.Duration] = duration.Milliseconds()
		logrus.WithFields(fields).Logf(level, format, args...)
	}
}

func Log(ctx context.Context, level logrus.Level, args ...interface{}) {
	if logrus.IsLevelEnabled(level) {
		fields := AppendFields(ctx, logrus.Fields{})
		if len(args) == 1 {
			msg := args[0]
			s1, ok := msg.(string)
			if ok {
				logrus.WithFields(fields).Log(level, s1)
			} else {
				bs, err := json.Marshal(msg)
				if err != nil {
					logrus.WithFields(fields).Log(level, args...)
				} else {
					s2 := string(bs)
					logrus.WithFields(fields).Log(level, s2)
				}
			}
		} else {
			logrus.WithFields(fields).Log(level, args...)
		}
	}
}
func Logf(ctx context.Context, level logrus.Level, format string, args ...interface{}) {
	fields := AppendFields(ctx, logrus.Fields{})
	logrus.WithFields(fields).Logf(level, format, args...)
}

func AppendFields(ctx context.Context, fields logrus.Fields) logrus.Fields {
	if len(fieldConfig.FieldMap) > 0 {
		if logFields, ok := ctx.Value(fieldConfig.FieldMap).(map[string]interface{}); ok {
			for k, v := range logFields {
				fields[k] = v
			}
		}
	}
	if fieldConfig.Fields != nil {
		cfs := *fieldConfig.Fields
		for _, k2 := range cfs {
			if v2, ok := ctx.Value(k2).(string); ok && len(v2) > 0 {
				fields[k2] = v2
			}
		}
	}
	return fields
}

func Panic(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.PanicLevel, args...)
}
func Fatal(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.FatalLevel, args...)
}
func Error(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.ErrorLevel, args...)
}
func Warn(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.WarnLevel, args...)
}
func Info(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.InfoLevel, args...)
}
func Debug(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.DebugLevel, args...)
}
func Trace(ctx context.Context, args ...interface{}) {
	Log(ctx, logrus.TraceLevel, args...)
}

func Panicf(ctx context.Context, format string, args ...interface{}) {
	Logf(ctx, logrus.PanicLevel, format, args...)
}
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	Logf(ctx, logrus.FatalLevel, format, args...)
}
func Errorf(ctx context.Context, format string, args ...interface{}) {
	Logf(ctx, logrus.ErrorLevel, format, args...)
}
func Warnf(ctx context.Context, format string, args ...interface{}) {
	Logf(ctx, logrus.WarnLevel, format, args...)
}
func Infof(ctx context.Context, format string, args ...interface{}) {
	Logf(ctx, logrus.InfoLevel, format, args...)
}
func Debugf(ctx context.Context, format string, args ...interface{}) {
	Logf(ctx, logrus.DebugLevel, format, args...)
}
func Tracef(ctx context.Context, format string, args ...interface{}) {
	Logf(ctx, logrus.TraceLevel, format, args...)
}
