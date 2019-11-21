package model

type Interface interface {
	GetId() int
	Validate() error
	TableName() string
	GetTableCacheKey() string
	GetItemCacheKey() string
}
