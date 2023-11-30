package wd_database

import (
	"github.com/boltdb/bolt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	path = "private.db"
)

func Open(dsn string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:                              false,
		QueryFields:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})

	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(3)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetMaxOpenConns(10)
	return db, err
}

var (
	BoltDB *bolt.DB
)

func OpenVoltDB(dsn string) *bolt.DB {
	bolt, err := bolt.Open(dsn, 0600, nil)
	if err != nil {
		log.Fatalf("can open volt db %+v", err)
	}
	return bolt
}

func ConnectPrivateVoltDB() *bolt.DB {
	if BoltDB == nil {
		BoltDB = OpenVoltDB("private.db")
	}
	return BoltDB
}
