package store

import (
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Store struct {
	Db    *dbx.DB
	Cache *Cache
}

func New(db *dbx.DB, cache *Cache) *Store {
	return &Store{Db: db, Cache: cache}
}
