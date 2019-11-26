package model

type Interface interface {
	GetId() uint
	Validate() error
	TableName() string
	GetTableCacheKey() string
	GetItemCacheKey() string
}
