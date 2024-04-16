package config

type Logger struct {
	Output string `mapstructure:"output"`
	Level  string `mapstructure:"level"`
}
