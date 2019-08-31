package model

import (
	"begin/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func SetUp() {
	dsn := setting.DatabaseSetting.Path
	db, err := gorm.Open("sqlite3", dsn)

	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
