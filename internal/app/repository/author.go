package repository

import (
	"encoding/json"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
)

type AuthorRepository struct {
	Store *store.Store
}

func (r AuthorRepository) Create(m *model.Author) error {
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

func (r AuthorRepository) Find(id uint) (*model.Author, error) {
	var m model.Author
	val, err := r.Store.Cache.Get(m.GetItemCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Find(&m, id).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetItemCacheKey(), m)
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

func (r AuthorRepository) FindAll() ([]*model.Author, error) {
	var res []*model.Author
	var m model.Author
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

func (r AuthorRepository) Update(m *model.Author) error {
	//TODO remake on callbacks
	if err := m.Validate(); err != nil {
		return err
	}
	err := r.Store.Db.Save(m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return err
}

func (r AuthorRepository) Delete(m model.Author) (err error) {
	err = r.Store.Db.Delete(&m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return
}

func (r AuthorRepository) GetOrdered(column string) ([]*model.Author, error) {
	var m model.Author
	var res []*model.Author
	val, err := r.Store.Cache.Get(m.GetTableCacheKey() + "_ordered")
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Order(column).Find(&res).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetTableCacheKey()+"_ordered", res)
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

func (r AuthorRepository) FindByLogin(login string) (*model.Author, error) {
	var m model.Author
	val, err := r.Store.Cache.Get(m.GetItemCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Where("login = ?", login).Find(&m).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetItemCacheKey(), m)
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
