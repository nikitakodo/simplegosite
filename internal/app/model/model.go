package model

import (
	"simplesite/internal/app/store"
)

type Interface interface {
	GetId() int
	Validate() error
	TableName() string
	GetTableCacheKey() string
	GetItemCacheKey() string
	GetCacheService() *store.Cache
}
