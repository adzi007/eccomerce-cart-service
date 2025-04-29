package database

import (
	"cart-service/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqliteDatabase struct {
	Db *gorm.DB
}

var sqliteInstance *sqliteDatabase

func NewSqliteDatabase() Database {

	dbName := config.ENV.DB_NAME

	if dbName == "" {
		dbName = "ecommerce-cart.db" // Default filename if DB_NAME is empty
	}

	// Ensure parent dir exists
	if err := os.MkdirAll(filepath.Dir(dbName), os.ModePerm); err != nil {
		log.Fatalf("Failed to create parent directory for SQLite DB: %v", err)
	}

	dsn := fmt.Sprintf("%s", dbName)

	var db *gorm.DB
	var err error

	// Retry connecting to the database up to 10 times
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to SQLite. Retrying in 2 seconds... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		panic("failed to connect to SQLite after 10 attempts")
	}

	sqliteInstance = &sqliteDatabase{Db: db}
	return sqliteInstance
}

func (p *sqliteDatabase) GetDb() *gorm.DB {
	return sqliteInstance.Db
}
