package repository

import (
	"encoding/json"
	"fmt"
	"simplesite/internal/app/model"
)

type NavRepository struct {
	Repository
}

func (r NavRepository) GetOrdered(m model.Interface) ([]model.Nav, error) {
	val, err := r.Store.Cache.Get(
		m.GetTableCacheKey() + "_ordered",
	)
	if err != nil {
		fmt.Println("01", err)
		return nil, err
	}
	var res []model.Nav
	if val == nil || *val == "" {
		err := r.Store.Db.Select().From(m.TableName()).OrderBy("order").All(res)
		if err != nil {
			fmt.Println("02", err)
			return nil, err
		}
		err = r.Store.Cache.SetStruct(
			m.GetTableCacheKey()+"_ordered",
			res,
		)
		if err != nil {
			fmt.Println("03", err)
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &res)
		if err != nil {
			fmt.Println("04", err)
			return nil, err
		}
	}
	return res, nil
}
