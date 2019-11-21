package repository

import (
	"simplesite/internal/app/store"
)

type AddRepository struct {
	Repository
	Store *store.Store
}
