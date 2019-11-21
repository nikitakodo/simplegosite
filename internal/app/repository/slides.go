package repository

import (
	"encoding/json"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
)

type SlidesRepository struct {
	Repository
	Store *store.Store
}

func (r *SlidesRepository) GetOrdered(m model.Interface) ([]model.Interface, error) {
	val, err := r.Store.Cache.Get(
		m.GetTableCacheKey() + "_ordered",
	)
	if err != nil {
		return nil, err
	}
	var res []model.Interface
	if val == nil || *val == "" {
		err := r.Store.Db.Select().From(m.TableName()).OrderBy("order").All(res)
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(
			m.GetTableCacheKey()+"_ordered",
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
