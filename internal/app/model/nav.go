package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
	"simplesite/internal/app/store"
	"strconv"
)

type Nav struct {
	Cache      *store.Cache
	ID         int64        `gorm:"primary_key"`
	Order      int          `gorm:"column:order"`
	Title      string       `gorm:"column:title"`
	Uri        string       `gorm:"column:uri"`
	CreateTime sql.NullTime `gorm:"column:create_time"`
	UpdateTime sql.NullTime `gorm:"column:update_time"`
}

func (m Nav) GetCacheService() *store.Cache {
	return m.Cache
}

func (m Nav) GetId() int {
	return m.ID
}

func (m Nav) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Nav) GetItemCacheKey() string {
	return m.TableName() + "_" + strconv.Itoa(m.GetId())
}

func (m Nav) TableName() string {
	return "nav"
}

func (m Nav) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Order, validation.Required),
		validation.Field(&m.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.Uri, validation.Required),
	)
}
