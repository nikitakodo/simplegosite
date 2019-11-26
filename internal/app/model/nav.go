package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Nav struct {
	gorm.Model
	Order int    `gorm:"column:order" json:"order"`
	Title string `gorm:"column:title" json:"title"`
	Uri   string `gorm:"column:uri" json:"uri"`
}

func (m Nav) GetId() uint {
	return m.ID
}

func (m Nav) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Nav) GetItemCacheKey() string {
	return m.TableName() + "_" + string(m.GetId())
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
