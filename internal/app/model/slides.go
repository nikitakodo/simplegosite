package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
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

func (s Slide) BeforeCreate(model *Interface) error {
	return nil
}

func (s Slide) AfterCreate(model *Interface) error {
	return nil
}

func (s Slide) BeforeUpdate(model *Interface) error {
	return nil
}

func (s Slide) AfterUpdate(model *Interface) error {
	return nil
}

func (s Slide) BeforeDelete(model *Interface) error {
	return nil
}

func (s Slide) AfterDelete(model *Interface) error {
	return nil
}

func (s Slide) TableName() string {
	return "slides"
}

func (s Slide) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.Order, validation.Required),
		validation.Field(&s.FirstTitle, validation.Required, validation.Length(6, 100)),
		validation.Field(&s.SecondTitle, validation.Required, validation.Length(6, 100)),
		validation.Field(&s.ThirdTitle, validation.Required, validation.Length(6, 100)),
		validation.Field(&s.Img, validation.Required),
	)
}
