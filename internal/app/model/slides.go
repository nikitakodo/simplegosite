package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

type Slide struct {
	gorm.Model
	Order       int    `json:"order"`
	FirstTitle  string `json:"first_title"`
	SecondTitle string `json:"second_title"`
	ThirdTitle  string `json:"third_title"`
	Img         string `json:"img"`
}

func (m Slide) GetId() uint {
	return m.ID
}

func (m Slide) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Slide) GetItemCacheKey() string {
	return m.TableName() + "_" + string(m.GetId())
}

func (m Slide) TableName() string {
	return "slides"
}

func (m Slide) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Order, validation.Required),
		validation.Field(&m.FirstTitle, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.SecondTitle, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.ThirdTitle, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.Img, validation.Required),
	)
}
