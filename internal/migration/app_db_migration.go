package main

import (
	"cart-service/config"
	"cart-service/config/database"
	"cart-service/internal/model/entity"
)

func main() {
	config.LoadConfig()

	db := database.NewMysqlDatabase()

	appDbMigrate(db)
}

func appDbMigrate(db database.Database) {

	// db.GetDb().Migrator().CreateTable(&entity.Cart{})

	err := db.GetDb().Migrator().AutoMigrate(&entity.Cart{})

	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

}
