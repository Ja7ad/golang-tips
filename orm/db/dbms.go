package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func SqliteConnector() error {
	DB, err = gorm.Open(sqlite.Open(os.Getenv("SQLITEPATH")), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate data to table
	if err := DB.AutoMigrate(&Books{}); err != nil {
		return err
	}
	return nil
}

func MysqlConnector() error {
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       os.Getenv("MYSQLCONFIG"),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate data to table
	if err := DB.AutoMigrate(&Books{}); err != nil {
		return err
	}
	return nil
}
