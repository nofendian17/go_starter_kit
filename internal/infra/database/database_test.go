package database

import (
	"github.com/nofendian17/gostarterkit/pkg/logger"
	"gorm.io/gorm"
	"testing"

	"github.com/nofendian17/gostarterkit/internal/config"
	mocks "github.com/nofendian17/gostarterkit/internal/mocks/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	l := logger.New(logger.Config{
		File: logger.File{
			IsActive: false,
			LogFile:  "/tmp/app.log",
			Format:   "json",
		},
		Console: logger.Console{
			Format: "text",
		},
	})
	oldGormOpen := gormOpen
	defer func() {
		gormOpen = oldGormOpen
	}()

	testCases := []struct {
		name   string
		driver string
	}{
		{name: "Postgres", driver: "postgres"},
		{name: "MySQL", driver: "mysql"},
		{name: "SQLServer", driver: "sqlserver"},
		{name: "SQLite", driver: "sqlite"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gormOpen = func(dialector gorm.Dialector, opts ...gorm.Option) (db *gorm.DB, err error) {
				DB, sqlMock, err := mocks.New(tc.driver)
				sqlMock.ExpectPing().WillReturnError(nil)
				if err != nil {
					t.Fatal(err)
				}
				return DB.GormDB, nil
			}

			cfg.Database.Driver = tc.driver
			db, err := New(cfg, l)
			if err != nil {
				assert.FailNow(t, err.Error())
			}
			defer db.SqlDB.Close()

			assert.IsType(t, &DB{}, db)
		})
	}
}
