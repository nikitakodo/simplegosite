package sqlstore

import (
	"database/sql"
	"simplesite/internal/app/store"
)

type Store struct {
	db *sql.DB
}

func (s *Store) GetRepository(repositoryName string) (store.RepositoryInterface, error) {
	switch repositoryName {
	case "admin":
		return AdminRepository{store: s}, nil
	}
	return nil, store.RepositoryNotFound
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}
