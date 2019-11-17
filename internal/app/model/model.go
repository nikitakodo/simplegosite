package model

type Interface interface {
	BeforeCreate(model *Interface) error
	AfterCreate(model *Interface) error
	BeforeUpdate(model *Interface) error
	AfterUpdate(model *Interface) error
	BeforeDelete(model *Interface) error
	AfterDelete(model *Interface) error
	Validate() error
	TableName() string
}
