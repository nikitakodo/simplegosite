package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
	"strconv"
)

type Nav struct {
	ID         int          `gorm:"primary_key" json:"id"`
	Order      int          `gorm:"column:order" json:"order"`
	Title      string       `gorm:"column:title" json:"title"`
	Uri        string       `gorm:"column:uri" json:"uri"`
	CreateTime sql.NullTime `gorm:"column:create_time" json:"create_time"`
	UpdateTime sql.NullTime `gorm:"column:update_time" json:"update_time"`
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
