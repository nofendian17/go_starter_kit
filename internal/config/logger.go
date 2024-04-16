package config

type Logger struct {
	File    FileLogger    `mapstructure:"file"`
	Console ConsoleLogger `mapstructure:"console"`
}

type FileLogger struct {
	IsActive bool   `mapstructure:"isActive"`
	LogFile  string `mapstructure:"logFile"`
	Format   string `mapstructure:"format"`
}

type ConsoleLogger struct {
	Format string `mapstructure:"format"`
}
