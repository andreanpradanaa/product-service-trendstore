package db

import (
	"fmt"
	"time"

	"github.com/andreanpradanaa/product-service-trendstore/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDBConnection(dbconf config.Database) (*gorm.DB, error) {
	host := dbconf.Host
	user := dbconf.Username
	password := dbconf.Password
	dbname := dbconf.DBName
	port := dbconf.Port

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbname, port)
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
