package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
	"strconv"
)

type Slide struct {
	ID          int          `json:"id"`
	Order       int          `json:"order"`
	FirstTitle  string       `json:"first_title"`
	SecondTitle string       `json:"second_title"`
	ThirdTitle  string       `json:"third_title"`
	Img         string       `json:"img"`
	CreateTime  sql.NullTime `json:"create_time"`
	UpdateTime  sql.NullTime `json:"update_time"`
}

func (m Slide) GetId() int {
	return m.ID
}

func (m Slide) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Slide) GetItemCacheKey() string {
	return m.TableName() + "_" + strconv.Itoa(m.GetId())
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
