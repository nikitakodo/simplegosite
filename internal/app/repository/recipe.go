package repository

import (
	"simplesite/internal/app/store"
)

type RecipeRepository struct {
	Repository
	Store *store.Store
}
