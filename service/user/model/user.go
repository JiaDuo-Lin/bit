package model

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
)

type User struct {
	ID   int64
	Name string
	Tags []string
}

func NewUser(id int64, name string, tags []string) *User {
	return &User{
		ID:   id,
		Name: name,
		Tags: tags,
	}
}

func (user *User) transform() *userInDB {
	tags, _ := json.Marshal(user.Tags)
	return &userInDB{
		ID:   user.ID,
		Name: user.Name,
		Tags: string(tags),
	}
}

func (user *User) IsRegistered() bool {
	log.Fatal(!db.NewRecord(user.transform()))
	return !db.NewRecord(user.transform())
}

func (user *User) Load() (err error) {
	userDB := loadById(user.ID)
	if userDB == nil {
		err = errors.Wrapf(err, "Can't find the user  in database")
		return
	}

	newUser := userDB.transform()
	user.Name = newUser.Name
	user.Tags = newUser.Tags
	return
}

func (user *User) Add() {
	if !user.IsRegistered() {
		user.transform().Add()
	}
}

func (user *User) hasTags() bool {
	return len(user.Tags) > 0
}

func loadById(id int64) (newUser *userInDB) {
	db.First(newUser, id)
	return
}
