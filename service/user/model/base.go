package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"user/config"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", config.Db["db1"].Dsn)
	if err != nil {
		panic("数据库无法加载")
	}
}
