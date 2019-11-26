package model

import (
	"github.com/jinzhu/gorm"
)

type Social struct {
	gorm.Model
	Order int    `json:"order" gorm:"column:order"`
	Icon  string `json:"icon" gorm:"column:icon"`
	Url   string `json:"url" gorm:"column:url"`
}

func (m Social) GetId() uint {
	return m.ID
}

func (m Social) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Social) GetItemCacheKey() string {
	return m.TableName() + "_" + string(m.GetId())
}

func (m Social) Validate() error {
	return nil
}

func (m Social) TableName() string {
	return "social"
}
