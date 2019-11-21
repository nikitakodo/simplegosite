package model

import (
	"database/sql"
	"simplesite/internal/app/store"
	"strconv"
)

type Social struct {
	Cache      *store.Cache
	ID         int          `json:"id"`
	Order      int          `json:"order"`
	Icon       string       `json:"icon"`
	Url        string       `json:"url"`
	CreateTime sql.NullTime `json:"create_time"`
	UpdateTime sql.NullTime `json:"update_time"`
}

func (m Social) GetCacheService() *store.Cache {
	return m.Cache
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
