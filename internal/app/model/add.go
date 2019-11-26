package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Add struct {
	gorm.Model
	Title      string
	FirstItem  string
	SecondItem string
	ThirdItem  string
	FourthItem string
	FirstImg   string
	SecondImg  string
	ThirdImg   string
}

func (m Add) GetId() uint {
	return m.ID
}

func (m Add) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Add) GetItemCacheKey() string {
	return m.TableName() + "_" + string(m.GetId())
}

func (m Add) TableName() string {
	return "add"
}

func (m Add) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.FirstItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.SecondItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.ThirdItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.FourthItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.FirstImg, validation.Required),
		validation.Field(&m.SecondImg, validation.Required),
		validation.Field(&m.ThirdImg, validation.Required),
	)
}
