package database

import (
	"cart-service/config"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *mysqlDatabase
)

func NewMysqlDatabase() Database {

	once.Do(func() {

		dbUsername := config.ENV.DB_USERNAME
		dbPassword := config.ENV.DB_PASSWORD
		dbName := config.ENV.DB_NAME
		dbHost := config.ENV.DB_HOST
		dbPort := config.ENV.DB_PORT

		connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

		db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
		})

		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = &mysqlDatabase{Db: db}
	})

	return dbInstance
}

func (p *mysqlDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
