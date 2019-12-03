package model

import "github.com/jinzhu/gorm"

type Settings struct {
	gorm.Model
	Key   string
	Value string
}

func (s Settings) GetId() uint {
	return s.ID
}

func (s Settings) Validate() error {
	return nil
}

func (s Settings) TableName() string {
	return "settings"
}

func (s Settings) GetTableCacheKey() string {
	return s.TableName() + "_all"
}

func (s Settings) GetItemCacheKey() string {
	return s.TableName() + "_" + string(s.GetId())
}
