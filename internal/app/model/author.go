package model

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Recipe   []Recipe `gorm:"foreignkey:AuthorId"`
	Name     string
	Login    string
	Password string
	IsBanned bool
}

func (a Author) GetId() uint {
	return a.ID
}

func (a Author) Validate() error {
	return nil
}

func (a Author) TableName() string {
	return "author"
}

func (a Author) GetTableCacheKey() string {
	return a.TableName() + "_all"
}

func (a Author) GetItemCacheKey() string {
	return a.TableName() + "_" + string(a.GetId())
}
