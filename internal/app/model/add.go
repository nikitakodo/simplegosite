package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Add struct {
	ID         int          `json:"id"`
	Title      string       `json:"title"`
	FirstItem  string       `json:"first_item"`
	SecondItem string       `json:"second_item"`
	ThirdItem  string       `json:"third_item"`
	FourthItem string       `json:"fourth_item"`
	FirstImg   string       `json:"first_img"`
	SecondImg  string       `json:"second_img"`
	ThirdImg   string       `json:"third_img"`
	CreateTime sql.NullTime `json:"create_time"`
	UpdateTime sql.NullTime `json:"update_time"`
}

func (a Add) BeforeCreate(model *Interface) error {
	return nil
}

func (a Add) AfterCreate(model *Interface) error {
	return nil
}

func (a Add) BeforeUpdate(model *Interface) error {
	return nil
}

func (a Add) AfterUpdate(model *Interface) error {
	return nil
}

func (a Add) BeforeDelete(model *Interface) error {
	return nil
}

func (a Add) AfterDelete(model *Interface) error {
	return nil
}

func (a Add) TableName() string {
	return "add"
}

func (a Add) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&a.FirstItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&a.SecondItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&a.ThirdItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&a.FourthItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&a.FirstImg, validation.Required),
		validation.Field(&a.SecondImg, validation.Required),
		validation.Field(&a.ThirdImg, validation.Required),
	)
}
