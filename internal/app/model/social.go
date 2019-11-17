package model

import "database/sql"

type Social struct {
	ID         int          `json:"id"`
	Order      int          `json:"order"`
	Icon       string       `json:"icon"`
	Url        string       `json:"url"`
	CreateTime sql.NullTime `json:"create_time"`
	UpdateTime sql.NullTime `json:"update_time"`
}

func (s Social) BeforeCreate(model *Interface) error {
	return nil
}

func (s Social) AfterCreate(model *Interface) error {
	return nil
}

func (s Social) BeforeUpdate(model *Interface) error {
	return nil
}

func (s Social) AfterUpdate(model *Interface) error {
	return nil
}

func (s Social) BeforeDelete(model *Interface) error {
	return nil
}

func (s Social) AfterDelete(model *Interface) error {
	return nil
}

func (s Social) Validate() error {
	return nil
}

func (s Social) TableName() string {
	return "social"
}
