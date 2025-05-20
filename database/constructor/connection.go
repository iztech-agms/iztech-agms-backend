package constructor

import (
	"fmt"
	"graduation-system/globals"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(dsn string) error {
	var err error
	globals.GMSDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error while connecting to the database: %v", err)
	}

	// Create database if not exists
	if err := globals.GMSDB.Exec("CREATE DATABASE IF NOT EXISTS gms").Error; err != nil {
		return fmt.Errorf("error while creating database: %v", err)
	}

	// Connection pool settings
	sqlDB, err := globals.GMSDB.DB()
	if err != nil {
		return fmt.Errorf("error while getting database connection pool: %v", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)      // Maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)     // Maximum number of open connections
	sqlDB.SetConnMaxLifetime(3600) // Maximum lifetime of connections (seconds)

	fmt.Println("Successfully connected to the database!")
	return nil
}
