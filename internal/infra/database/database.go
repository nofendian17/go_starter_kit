package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/nofendian17/gostarterkit/internal/config"
	l "github.com/nofendian17/gostarterkit/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	GormDB *gorm.DB
	SqlDB  *sql.DB
}

var gormOpen = gorm.Open

func New(cfg *config.Config, l l.Logger) (*DB, error) {
	ctx := context.Background()
	debugMode := logger.Silent
	if cfg.Database.Debug {
		debugMode = logger.Info
	}

	dialect, err := getDialect(cfg)
	if err != nil {
		l.Error(ctx, "Failed to create database dialect", err)
		return nil, err
	}
	l.Info(ctx, fmt.Sprintf("Connecting to database %s with driver %s", cfg.Database.Database, cfg.Database.Driver), nil)

	gormDB, err := gormOpen(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(debugMode),
	})
	if err != nil {
		l.Error(ctx, "Failed to connect to database", err)
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		l.Error(ctx, "Failed to get SQL DB", err)

		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)

	if err = sqlDB.Ping(); err != nil {
		l.Error(ctx, "Failed to ping database: %v", err)
		return nil, err
	}
	l.Info(ctx, fmt.Sprintf("Successfully connected to database %s with driver %s", cfg.Database.Database, cfg.Database.Driver), nil)

	return &DB{
		GormDB: gormDB,
		SqlDB:  sqlDB,
	}, nil
}

func getDialect(cfg *config.Config) (gorm.Dialector, error) {
	switch cfg.Database.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database, cfg.Database.Charset)
		return mysql.Open(dsn), nil

	case "postgres":
		sslMode := "disable"
		if cfg.Database.SSLMode {
			sslMode = "enable"
		}
		dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s TimeZone=%s",
			cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Database, cfg.Database.Password, sslMode, cfg.Database.Timezone)
		return postgres.Open(dsn), nil

	case "sqlite":
		return sqlite.Open(cfg.Database.Database), nil

	case "sqlserver":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
			cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Database)
		return sqlserver.Open(dsn), nil

	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}
}
