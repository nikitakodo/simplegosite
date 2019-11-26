package model

import "github.com/jinzhu/gorm"

type Cuisine struct {
	gorm.Model
	Recipe      []Recipe `gorm:"foreignkey:CuisineId"`
	Name        string
	Description string
}

func (c Cuisine) GetId() uint {
	return c.ID
}

func (c Cuisine) Validate() error {
	return nil
}

func (c Cuisine) TableName() string {
	return "cuisine"
}

func (c Cuisine) GetTableCacheKey() string {
	return c.TableName() + "_all"
}

func (c Cuisine) GetItemCacheKey() string {
	return c.TableName() + "_" + string(c.GetId())
}
