package db

import (
	"fmt"
	"vkbot/internal/users"
)

type Entity interface{}

type DB interface {
	Get(id int64) (Entity, error)
	Add(id int64, e Entity) error
	Delete(id int64) error
	Exists(id int64) (bool, error)
}

func NewDB() DB {
	return &MapDB{
		data: map[int64]users.User{},
	}
}

type MapDB struct {
	data map[int64]users.User
}

func (db *MapDB) Get(id int64) (Entity, error) {
	if exists, _ := db.Exists(id); !exists {
		return nil, fmt.Errorf("can't get user, user doesn't exist")
	}
	return db.data[id], nil
}

func (db *MapDB) Add(id int64, e Entity) error {
	db.data[id] = e.(users.User)
	return nil
}

func (db *MapDB) Delete(id int64) error {
	delete(db.data, id)
	return nil
}

func (db *MapDB) Exists(id int64) (bool, error) {
	_, ok := db.data[id]
	return ok, nil
}
