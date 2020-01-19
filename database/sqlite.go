package database

import (
	"github.com/bonggar/gorestapi/config"
	"github.com/bonggar/gorestapi/model"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //sqlite3
)

//SQLiteDBConnect : Create Connection to database
func SQLiteDBConnect() {
	//Connect to database, exit when errored
	db, err = gorm.Open("sqlite3", "./"+config.DbName+".db")
	if err != nil {
		panic(err)
	}

	//If set true then print all executed queries to the console
	db.LogMode(config.DbDebug)

	//Migrate database
	SQLiteMigrate(db)
}

//SQLiteMigrate : do auto migration
func SQLiteMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&model.User{}).Error; err != nil {
		panic("Failed migrating database")
	}
}
