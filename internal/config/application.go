package config

type Application struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}
