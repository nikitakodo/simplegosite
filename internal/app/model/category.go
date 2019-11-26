package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Recipe      []Recipe `gorm:"foreignkey:CategoryId"`
	Name        string
	Description string
}

func (c Category) GetId() uint {
	return c.ID
}

func (c Category) Validate() error {
	return nil
}

func (c Category) TableName() string {
	return "category"
}

func (c Category) GetTableCacheKey() string {
	return c.TableName() + "_all"
}

func (c Category) GetItemCacheKey() string {
	return c.TableName() + "_" + string(c.GetId())
}
