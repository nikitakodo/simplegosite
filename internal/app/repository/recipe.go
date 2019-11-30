package repository

import (
	"encoding/json"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
	"strconv"
)

type RecipeRepository struct {
	Store *store.Store
}

func (r RecipeRepository) Create(m *model.Recipe) error {
	//TODO remake on callbacks
	if err := m.Validate(); err != nil {
		return err
	}
	err := r.Store.Db.Create(m).Error
	if err != nil {
		return err
	}
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	if err != nil {
		return err
	}
	return nil
}

func (r RecipeRepository) Find(id uint) (*model.Recipe, error) {
	var m model.Recipe
	val, err := r.Store.Cache.Get(m.GetItemCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Preload("Author").
			Preload("Category").
			Preload("Cuisine").
			Preload("Mark").Find(&m, id).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetTableCacheKey(), m)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &m)
		if err != nil {
			return nil, err
		}
	}
	return &m, nil
}

func (r RecipeRepository) FindAll() ([]*model.Recipe, error) {
	var res []*model.Recipe
	var m model.Recipe
	val, err := r.Store.Cache.Get(m.GetTableCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Find(&res).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetTableCacheKey(), res)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (r RecipeRepository) Update(m *model.Recipe) error {
	//TODO remake on callbacks
	if err := m.Validate(); err != nil {
		return err
	}
	err := r.Store.Db.Save(m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return err
}

func (r RecipeRepository) Delete(m model.Recipe) (err error) {
	err = r.Store.Db.Delete(&m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return
}

func (r RecipeRepository) GetLatest(limit int, offset int) ([]*model.Recipe, error) {
	var res []*model.Recipe
	var m model.Recipe
	val, err := r.Store.Cache.Get(
		m.GetTableCacheKey() + "_latest_l" + strconv.Itoa(limit) + "_o" + strconv.Itoa(offset),
	)
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.
			Order("created_at DESC").
			Preload("Author").
			Preload("Category").
			Preload("Cuisine").
			Preload("Mark").
			Offset(offset).
			Limit(limit).
			Find(&res).
			Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(
			m.GetTableCacheKey()+"_latest_l"+strconv.Itoa(limit)+"_o"+strconv.Itoa(offset),
			res,
		)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (r RecipeRepository) TopRated(limit int, offset int) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	var m model.Recipe
	val, err := r.Store.Cache.Get(
		m.GetTableCacheKey() + "_top_rated_l" + strconv.Itoa(limit) + "_o" + strconv.Itoa(offset),
	)
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {

		rows, err := r.Store.Db.Raw(`select r.* from recipe as r
left join mark m on r.id = m.recipe_id
where r.deleted_at is null
group by r.id,m.recipe_id
order by  count(m.id) desc`).Limit(limit).Offset(offset).Rows()
		defer rows.Close()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var recipe model.Recipe
			err = r.Store.Db.ScanRows(rows, &recipe)
			if err != nil {
				return nil, err
			}
			res, err := r.Find(recipe.GetId())
			if err != nil {
				return nil, err
			}
			recipes = append(recipes, res)
		}
		err = r.Store.Cache.SetStruct(
			m.GetTableCacheKey()+"_top_rated_l"+strconv.Itoa(limit)+"_o"+strconv.Itoa(offset),
			recipes,
		)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &recipes)
		if err != nil {
			return nil, err
		}
	}
	return recipes, nil
}

func (r RecipeRepository) MostLiked(limit int, offset int) ([]*model.Recipe, error) {
	var recipes []*model.Recipe
	var m model.Recipe
	val, err := r.Store.Cache.Get(
		m.GetTableCacheKey() + "_most_liked_l" + strconv.Itoa(limit) + "_o" + strconv.Itoa(offset),
	)
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {

		rows, err := r.Store.Db.Raw(`select r.* from recipe as r
join mark m on r.id = m.recipe_id
where r.deleted_at is null
group by r.id
order by  count(m.id) desc`).Limit(limit).Offset(offset).Rows()
		defer rows.Close()
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var recipe model.Recipe
			err = r.Store.Db.ScanRows(rows, &recipe)
			if err != nil {
				return nil, err
			}
			res, err := r.Find(recipe.GetId())
			if err != nil {
				return nil, err
			}
			recipes = append(recipes, res)
		}
		err = r.Store.Cache.SetStruct(
			m.GetTableCacheKey()+"_most_liked_l"+strconv.Itoa(limit)+"_o"+strconv.Itoa(offset),
			recipes,
		)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &recipes)
		if err != nil {
			return nil, err
		}
	}
	return recipes, nil
}
