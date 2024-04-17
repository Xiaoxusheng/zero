package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type DB struct {
	*gorm.DB
}

type Config struct {
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int64
}

func NewMysql(c *Config) (*DB, error) {
	db, err := gorm.Open(mysql.Open(c.Dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 nil,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(c.ConnMaxLifetime))

	return &DB{DB: db}, nil
}
