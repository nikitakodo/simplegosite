package repository

import (
	"database/sql"
	"encoding/json"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
)

type Interface interface {
	Create(abstractModel model.Interface) (*model.Interface, error)
	Find(abstractModel model.Interface) (*model.Interface, error)
	FindAll(abstractModel model.Interface) ([]model.Interface, error)
	Update(abstractModel model.Interface) (*model.Interface, error)
	Delete(abstractModel model.Interface) error
}

type Repository struct {
	Store *store.Store
}

func (r Repository) Create(m model.Interface) (*model.Interface, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	err := r.Store.Db.Model(m).Insert()
	if err != nil {
		return nil, err
	}
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r Repository) Find(m model.Interface) (*model.Interface, error) {
	val, err := r.Store.Cache.Get(m.GetItemCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Select().Model(m.GetId(), m)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetTableCacheKey(), m)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), m)
		if err != nil {
			return nil, err
		}
	}
	return &m, nil
}

func (r Repository) FindAll(m model.Interface) ([]model.Interface, error) {
	val, err := r.Store.Cache.Get(m.GetTableCacheKey())
	if err != nil {
		return nil, err
	}
	var res []model.Interface
	if val == nil || *val == "" {
		err := r.Store.Db.Select().All(res)
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

func (r Repository) Update(m model.Interface) (*model.Interface, error) {
	if err := m.Validate(); err != nil {
		return nil, err
	}
	err := r.Store.Db.Model(m).Update()
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return &m, err
}

func (r Repository) Delete(m model.Interface) (err error) {
	err = r.Store.Db.Model(m).Delete()
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return
}
