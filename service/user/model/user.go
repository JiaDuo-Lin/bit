package model

import (
	"github.com/2bitlab/bit_infra/conn"
	"github.com/2bitlab/bit_infra/logger"
	"github.com/pkg/errors"
	"context"
)

type User struct {
	ID    int64  `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100)"`
	Tags  []string  `gorm:"many2many:user_tangs";"ForeignKey:UserId"`
}


func NewUser(id int64, name, email string, tags []string) *User {

	return &User{
		ID:   id,
		Name: name,
		Email: email,
		Tags: tags,
	}
}


func (user *User) IsRegistered() bool {
	return !db.NewRecord(user)
}

func (user *User) Load() (err error) {
	userDB, err := loadById(user.ID)
	if err != nil {
		logger.TransLogger.Sugar().Errorf("Load has err[%v]", err)
		return
	}
	if userDB == nil {
		err = errors.Wrapf(err, "Can't find the user in database")
		logger.TransLogger.Sugar().Errorf("Load has err[%v]", err)
		return
	}
	return
}

func (user *User) Add() {
	if !user.IsRegistered() {
		db.Create(user)
	}
}

func (user *User) hasTags() bool {gi
	return len(user.Tags) > 0
}

func loadById(id int64) (newUser *User, err error) {
	db, err := conn.MySQLConn(context.Background())
	if err != nil {
		logger.TransLogger.Sugar().Errorf("MySQLConn has err[%v]", err)
		return
	}
	db.First(newUser, id)
	return
}
