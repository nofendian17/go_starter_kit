package config

type Database struct {
	Driver        string `mapstructure:"driver"`
	Host          string `mapstructure:"host"`
	Port          int    `mapstructure:"port"`
	Username      string `mapstructure:"username"`
	Password      string `mapstructure:"password"`
	Database      string `mapstructure:"database"`
	Charset       string `mapstructure:"charset"`
	Timezone      string `mapstructure:"timezone"`
	MaxIdleConns  int    `mapstructure:"maxIdleConns"`
	MaxOpenConns  int    `mapstructure:"maxOpenConns"`
	SSLMode       bool   `mapstructure:"sslMode"`
	Debug         bool   `mapstructure:"debug"`
	DisableMetric bool   `mapstructure:"disableMetric"`
	DisableTrace  bool   `mapstructure:"disableTrace"`
}
