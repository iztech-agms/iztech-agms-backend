package constructor

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error while connecting to the database: %v", err)
	}

	// Connection pool settings
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error while getting database connection pool: %v", err)
	}

	defer sqlDB.Close()

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)      // Maximum number of idle connections
	sqlDB.SetMaxOpenConns(100)     // Maximum number of open connections
	sqlDB.SetConnMaxLifetime(3600) // Maximum lifetime of connections (seconds)

	fmt.Println("Successfully connected to the database!")
	return nil
}
