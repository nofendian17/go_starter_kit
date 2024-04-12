package mocks

import (
	"database/sql"
	"errors"

	"github.com/DATA-DOG/go-sqlmock"
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

func New(driver string) (*DB, sqlmock.Sqlmock, error) {
	sqlDB, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		return nil, nil, err
	}

	var dialector gorm.Dialector
	switch driver {
	case "postgres":
		dialector = postgres.New(postgres.Config{
			Conn:       sqlDB,
			DriverName: "postgres",
		})
	case "mysql":
		dialector = mysql.New(mysql.Config{
			Conn:       sqlDB,
			DriverName: "mysql",
		})
		sqlMock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0"))

	case "sqlserver":
		dialector = sqlserver.New(sqlserver.Config{
			Conn:       sqlDB,
			DriverName: "sqlserver",
		})
	case "sqlite":
		dialector = sqlite.Open("file::memory:?cache=shared")
	default:
		sqlDB.Close()
		return nil, nil, errors.New("unsupported database driver")
	}

	sqlMock.ExpectPing()
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		sqlDB.Close()
		return nil, nil, err
	}

	return &DB{
		GormDB: db,
		SqlDB:  sqlDB,
	}, sqlMock, nil
}
