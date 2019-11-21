package repository

import (
	"simplesite/internal/app/store"
)

type SocialRepository struct {
	Repository
	Store *store.Store
}
