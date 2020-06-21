package model

import (
	"context"
	"github.com/2bitlab/bit_infra/conn"
	"github.com/2bitlab/bit_infra/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = conn.MySQLConn(context.TODO())
	if err != nil {
		logger.TransLogger.Sugar().Panic("MySQLConn has err[%v]", err)
	}
	return
}
