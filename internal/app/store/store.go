package store

import (
	"github.com/jinzhu/gorm"
)

type Store struct {
	Db    *gorm.DB
	Cache *Cache
}

func New(db *gorm.DB, cache *Cache) *Store {
	return &Store{Db: db, Cache: cache}
}
