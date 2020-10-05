package log

import "github.com/sirupsen/logrus"

type Config struct {
	Level           string           `mapstructure:"level"`
	TimestampFormat string           `mapstructure:"timestamp_format"`
	FieldMap        *logrus.FieldMap `mapstructure:"field_map"`
	Duration        string           `mapstructure:"duration"`
	Fields          string           `mapstructure:"fields"`
}

type FieldConfig struct {
	Duration string    `mapstructure:"duration"`
	Fields   *[]string `mapstructure:"fields"`
}
