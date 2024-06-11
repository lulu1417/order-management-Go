package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB
var err error

func InitDB() *gorm.DB {
	DBURL := os.Getenv("DB_URL")
	db, err = gorm.Open("mysql", DBURL)
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
