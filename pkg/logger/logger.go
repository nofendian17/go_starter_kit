package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

type Config struct {
	File    File
	Console Console
}

type File struct {
	IsActive bool
	LogFile  string
	Format   string
}

type Console struct {
	Format string
}

// Logger defines the interface for logging operations.
type Logger interface {
	Info(msg string, fields map[string]interface{})
	Warn(msg string, fields map[string]interface{})
	Error(msg string, fields map[string]interface{})
	Debug(msg string, fields map[string]interface{})
}

// logger implements the Logger interface.
type logger struct {
	log *slog.Logger
}

// Info logs an informational message with optional fields.
func (l *logger) Info(msg string, fields map[string]interface{}) {
	l.log.WithFields(fields).Info(msg)
	l.log.ResetExitHandlers()
}

// Warn logs a warning message with optional fields.
func (l *logger) Warn(msg string, fields map[string]interface{}) {
	l.log.WithFields(fields).Warn(msg)
}

// Error logs an error message with optional fields.
func (l *logger) Error(msg string, fields map[string]interface{}) {
	l.log.WithFields(fields).Error(msg)
}

// Debug logs a debug message with optional fields.
func (l *logger) Debug(msg string, fields map[string]interface{}) {
	l.log.WithFields(fields).Debug(msg)
}

// New creates a new Logger instance with default settings.
func New(cfg Config) Logger {
	fileHandler := handler.MustRotateFile(cfg.File.LogFile, rotatefile.EveryDay, handler.WithLogLevels(slog.AllLevels))
	switch cfg.File.Format {
	case "text":
		fileHandler.SetFormatter(slog.NewTextFormatter())
	default:
		fileHandler.SetFormatter(slog.NewJSONFormatter())
	}

	consoleHandler := handler.NewConsoleHandler(slog.AllLevels)
	switch cfg.Console.Format {
	case "json":
		consoleHandler.SetFormatter(slog.NewJSONFormatter())
	}

	log := slog.New()

	if cfg.File.IsActive {
		log.AddHandlers(consoleHandler, fileHandler)
	} else {
		log.AddHandler(consoleHandler)
	}

	return &logger{
		log: log,
	}
}
