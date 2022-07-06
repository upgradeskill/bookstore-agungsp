package system

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Storage *gorm.DB
}

func (db *Database) Init() {

	var err error 
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db.Storage, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}