package model

type File struct {
	Id         int64
	Name       string
	Path       string
	Size       int64
	Extension  string
	CreateTime string
	UpdateTime string
}

func (f File) BeforeCreate(model *Interface) error {
	return nil
}

func (f File) AfterCreate(model *Interface) error {
	return nil
}

func (f File) BeforeUpdate(model *Interface) error {
	return nil
}

func (f File) AfterUpdate(model *Interface) error {
	return nil
}

func (f File) BeforeDelete(model *Interface) error {
	return nil
}

func (f File) AfterDelete(model *Interface) error {
	return nil
}

func (f File) TableName() string {
	return "files"
}
