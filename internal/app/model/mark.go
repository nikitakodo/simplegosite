package model

import "github.com/jinzhu/gorm"

type Mark struct {
	gorm.Model
	RecipeId uint
	AuthorId uint
	Value    int
	Recipe   Recipe `gorm:"foreignkey:RecipeId"`
	Author   Author `gorm:"foreignkey:AuthorId"`
}

func (m Mark) GetId() uint {
	return m.ID
}

func (m Mark) Validate() error {
	return nil
}

func (m Mark) TableName() string {
	return "mark"
}

func (m Mark) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Mark) GetItemCacheKey() string {
	return m.TableName() + "_" + string(m.GetId())
}
