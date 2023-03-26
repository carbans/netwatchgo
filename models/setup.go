package models

import (
	"github.com/carbans/netdebug/config"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config config.Config) {
	database, err := gorm.Open(sqlite.Open(config.DBSource))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Agent{}, &User{}, &ApiKey{})

	DB = database
}
