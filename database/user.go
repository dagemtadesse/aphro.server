package database

import (
	"aphro.web/model"
)

func (db *Database) InsertUser(user *model.User) error {
	return db.Create(user).Error
}

func (db *Database) FetchUser(email string) (*model.User, error) {
	user := new(model.User)

	db.Where("email = ? ", email).Find(user)
	return user, db.Error
}
