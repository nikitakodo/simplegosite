package model

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
	"simplesite/internal/app/store"
	"strconv"
)

type Add struct {
	Cache      *store.Cache
	ID         int          `json:"id"`
	Title      string       `json:"title"`
	FirstItem  string       `json:"first_item"`
	SecondItem string       `json:"second_item"`
	ThirdItem  string       `json:"third_item"`
	FourthItem string       `json:"fourth_item"`
	FirstImg   string       `json:"first_img"`
	SecondImg  string       `json:"second_img"`
	ThirdImg   string       `json:"third_img"`
	CreateTime sql.NullTime `json:"create_time"`
	UpdateTime sql.NullTime `json:"update_time"`
}

func (m Add) GetCacheService() *store.Cache {
	return m.Cache
}

func (m Add) GetId() int {
	return m.ID
}

func (m Add) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Add) GetItemCacheKey() string {
	return m.TableName() + "_" + strconv.Itoa(m.GetId())
}

func (m Add) TableName() string {
	return "add"
}

func (m Add) Validate() error {
	return validation.ValidateStruct(
		m,
		validation.Field(&m.Title, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.FirstItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.SecondItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.ThirdItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.FourthItem, validation.Required, validation.Length(6, 100)),
		validation.Field(&m.FirstImg, validation.Required),
		validation.Field(&m.SecondImg, validation.Required),
		validation.Field(&m.ThirdImg, validation.Required),
	)
}
