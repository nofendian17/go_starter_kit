package database

import (
	"gorm.io/gorm"
	"testing"

	"github.com/nofendian17/gostarterkit/internal/config"
	mocks "github.com/nofendian17/gostarterkit/internal/mocks/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := config.New()
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
				DB, _, err := mocks.New(tc.driver)
				if err != nil {
					t.Fatal(err)
				}
				return DB.GormDB, nil
			}

			cfg.Database.Driver = tc.driver
			db := New(cfg)
			defer db.SqlDB.Close()

			assert.IsType(t, &DB{}, db)
		})
	}
}
