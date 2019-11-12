package store

import "errors"

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound  = errors.New("record not found")
	RepositoryNotFound = errors.New("repository not found")
)
