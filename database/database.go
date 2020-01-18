package database

import "github.com/jinzhu/gorm"

var db *gorm.DB
var err error

//GetDB : get current connection
func GetDB() *gorm.DB {

	return db
}
