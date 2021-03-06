package model

import (
	"github.com/jinzhu/gorm"
)

type Recipe struct {
	gorm.Model
	Title      string
	Body       string
	Banner     string
	Img        string
	CategoryId uint
	CuisineId  uint
	AuthorId   uint
	Mark       []Mark   `gorm:"foreignkey:RecipeId"`
	Category   Category `gorm:"foreignkey:CategoryId"`
	Cuisine    Cuisine  `gorm:"foreignkey:CuisineId"`
	Author     Author   `gorm:"foreignkey:AuthorId"`
	MarksCount int
}

func (m Recipe) GetId() uint {
	return m.ID
}

func (m Recipe) GetTableCacheKey() string {
	return m.TableName() + "_all"
}

func (m Recipe) GetItemCacheKey() string {
	return m.TableName() + "_" + string(m.GetId())
}

func (m Recipe) TableName() string {
	return "recipe"
}

func (m Recipe) Validate() error {
	return nil
}

func (m Recipe) Marks() int {
	var marks int
	d := len(m.Mark)
	if d == 0 {
		return d
	}
	for _, row := range m.Mark {
		marks += row.Value
	}
	return marks / d
}

func (m Recipe) MarksSlice() []bool {
	a := m.Marks()
	var s []bool = []bool{false, false, false, false, false}
	for i := 0; i < a; i++ {
		s[i] = true
	}
	return s
}

func (m Recipe) FormatedDate() string {
	return m.CreatedAt.Format("January 02, 2006")
}
