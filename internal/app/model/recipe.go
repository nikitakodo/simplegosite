package model

import (
	"database/sql"
	"simplesite/internal/app/store"
	"strconv"
)

type Recipe struct {
	Cache      *store.Cache
	ID         int          `json:"id"`
	Title      string       `json:"title"`
	Body       string       `json:"body"`
	Img        string       `json:"img"`
	CategoryId int          `json:"category_id"`
	CuisineId  int          `json:"cuisine_id"`
	AuthorId   int          `json:"author_id"`
	CreateTime sql.NullTime `json:"create_time"`
	UpdateTime sql.NullTime `json:"update_time"`
}

func (m Recipe) GetCacheService() *store.Cache {
	return m.Cache
}

func (m Recipe) GetId() int {
	return m.ID
}

func (m Recipe) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Recipe) GetItemCacheKey() string {
	return m.TableName() + "_" + strconv.Itoa(m.GetId())
}

func (m Recipe) TableName() string {
	return "recipe"
}

func (m Recipe) Validate() error {
	return nil
}
