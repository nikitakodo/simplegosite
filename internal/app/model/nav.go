package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Nav struct {
	ID    int    `json:"id"`
	Order int    `json:"order"`
	Title string `json:"title"`
	Uri   string `json:"uri"`
}

func (n Nav) BeforeCreate(model *Interface) error {
	return nil
}

func (n Nav) AfterCreate(model *Interface) error {
	return nil
}

func (n Nav) BeforeUpdate(model *Interface) error {
	return nil
}

func (n Nav) AfterUpdate(model *Interface) error {
	return nil
}

func (n Nav) BeforeDelete(model *Interface) error {
	return nil
}

func (n Nav) AfterDelete(model *Interface) error {
	return nil
}

func (n Nav) TableName() string {
	return "nav"
}

func (n Nav) Validate() error {
	return validation.ValidateStruct(
		n,
		validation.Field(&n.Order, validation.Required),
		validation.Field(&n.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&n.Uri, validation.Required),
	)
}
