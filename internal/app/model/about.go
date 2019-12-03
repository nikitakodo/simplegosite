package model

import "github.com/jinzhu/gorm"

type About struct {
	gorm.Model
	Title string
	Text  string
	Img   string
}

func (a About) GetId() uint {
	return a.ID
}

func (a About) Validate() error {
	return nil
}

func (a About) TableName() string {
	return "about"
}

func (a About) GetTableCacheKey() string {
	return a.TableName() + "_all"
}

func (a About) GetItemCacheKey() string {
	return a.TableName() + "_" + string(a.GetId())
}
