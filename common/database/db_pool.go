package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type DBPool map[string]*gorm.DB

func (pool *DBPool) ConnectUser() error {
	return nil
}
func (pool *DBPool) ConnectGroup() error {
	return nil
}
func (pool *DBPool) ConnectArea(dsn, logLevel string, maxIdle, macConn, maxLifeCycle int) (*gorm.DB, error) {
	DBLogger := getDBLogger(logLevel)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      DBLogger,
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	(*pool)[DBAREA] = db
	return db, nil
}
func (pool *DBPool) ConnectWorkout() error {
	return nil
}
func (pool *DBPool) ConnectDatabases(dsn, logLevel string, maxIdle, maxConn, maxLifeCycle int) (*gorm.DB, error) {
	DBLogger := getDBLogger(logLevel)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      DBLogger,
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := database.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifeCycle) * time.Second)
	sqlDB.SetMaxOpenConns(maxConn)

	return database, nil
}
func (pool *DBPool) ConnectAll(cfg *DBConfig) error {
	for k, v := range cfg.Databases {
		if db, err := pool.ConnectDatabases(v, cfg.LogLevel, cfg.MaxIdle, cfg.MaxConn, cfg.MaxLifeCycle); err != nil {
			logrus.WithField("database", k).Errorf("connect database failure - %s", err)
			return err
		} else {
			(*pool)[k] = db
		}
	}
	return nil
}

func getLevel(level string) logger.LogLevel {
	switch level {
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}
}
func getDBLogger(level string) logger.Interface {
	return logger.New(
		log.New(logrus.StandardLogger().Out, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      getLevel(level),
			Colorful:      true,
		})
}

func NewDBPool() DBPool {
	return make(DBPool)
}
