package log

import "github.com/sirupsen/logrus"

type Config struct {
	Level           string           `mapstructure:"level"`
	TimestampFormat string           `mapstructure:"timestamp_format"`
	Map             *logrus.FieldMap `mapstructure:"map"`
	Duration        string           `mapstructure:"duration"`
	Fields          string           `mapstructure:"fields"`
	FieldMap        string           `mapstructure:"field_map"`
}

type FieldConfig struct {
	FieldMap string    `mapstructure:"field_map"`
	Duration string    `mapstructure:"duration"`
	Fields   *[]string `mapstructure:"fields"`
}
