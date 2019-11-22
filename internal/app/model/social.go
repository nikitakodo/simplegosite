package model

import (
	"database/sql"
	"strconv"
)

type Social struct {
	ID         int          `json:"id"`
	Order      int          `json:"order" gorm:"column:order"`
	Icon       string       `json:"icon" gorm:"column:icon"`
	Url        string       `json:"url" gorm:"column:url"`
	CreateTime sql.NullTime `json:"create_time" gorm:"column:create_time"`
	UpdateTime sql.NullTime `json:"update_time" gorm:"column:update_time"`
}

func (m Social) GetId() int {
	return m.ID
}

func (m Social) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Social) GetItemCacheKey() string {
	return m.TableName() + "_" + strconv.Itoa(m.GetId())
}

func (m Social) Validate() error {
	return nil
}

func (m Social) TableName() string {
	return "social"
}
