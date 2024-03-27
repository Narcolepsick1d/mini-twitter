package database

import (
	"Narcolepsick1d/mini-twitter/internal/config"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // dialect
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewSQLDB returns raw sql.DB entity.
func NewSQLDB(config *config.Config) (*sql.DB, error) {
	if config == nil {
		return nil, errors.New("missing config")
	}
	dsn := config.Database.DSN()
	if config.Database.SSLMode == "" {
		dsn += " sslmode=disable"
	} else {
		dsn += fmt.Sprintf(" sslmode=%s", config.Database.SSLMode)
	}
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	if config.Database.MaxConns != 0 {
		sqlDB.SetMaxOpenConns(config.Database.MaxConns)
	} else {
		sqlDB.SetMaxOpenConns(15)
	}

	// SetMaxIdleConns sets the maximum number of idle connections to the database.
	if config.Database.MaxIdleConnections != 0 {
		sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnections)
	} else {
		sqlDB.SetMaxIdleConns(10)
	}

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	if config.Database.MaxConnLifeTimeInSeconds != 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.Database.MaxConnLifeTimeInSeconds))
	} else {
		sqlDB.SetConnMaxLifetime(25 * time.Minute)
	}

	// SetConnMaxIdletime sets the maximum amount of time an idle connection may be reused.
	if config.Database.MaxConnIdleTimeInSeconds != 0 {
		sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(config.Database.MaxConnIdleTimeInSeconds))
	} else {
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	}

	return sqlDB, nil
}

// NewGoquDB returns raw goqu.Database entity.
func NewGoquDB(config *config.Config) (*goqu.Database, error) {
	sqlDB, err := NewSQLDB(config)
	if err != nil {
		return nil, err
	}

	dialect := goqu.Dialect("postgres")
	db := dialect.DB(sqlDB)

	return db, nil
}
