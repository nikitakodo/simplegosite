package sqlstore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"simplesite/internal/app/model"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
	"strconv"
)

type AddRepository struct {
	Store *Store
	Cache *services.Cache
}

func (repo *AddRepository) Create(add *model.Add) error {

	if err := add.Validate(); err != nil {
		return err
	}

	err := repo.Store.Db.QueryRow(
		fmt.Sprintf(
			"INSERT INTO %s (title, first_item, second_item, third_item, fourth_item, first_img, second_img, third_img) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
			add.TableName(),
		),
		add.Title,
		add.FirstItem,
		add.SecondItem,
		add.ThirdItem,
		add.FourthItem,
		add.FirstImg,
		add.SecondImg,
		add.ThirdImg,
	).Scan(add.ID)

	if err != nil {
		return err
	}

	err = repo.Cache.Del(add.TableName() + "_all")
	if err != nil {
		return err
	}

	return nil
}

func (repo *AddRepository) Find(id int) (*model.Add, error) {
	add := &model.Add{}

	val, err := repo.Cache.Get(add.TableName() + "_" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	if val == nil || *val == "" {
		if err := repo.Store.Db.QueryRow(
			fmt.Sprintf(
				"SELECT id, title, first_item, second_item, third_item, fourth_item, first_img, second_img, third_img FROM %s WHERE id = $1",
				add.TableName(),
			),
			id,
		).Scan(
			&add.ID,
			&add.Title,
			&add.FirstItem,
			&add.SecondItem,
			&add.ThirdItem,
			&add.FourthItem,
			&add.FirstImg,
			&add.SecondImg,
			&add.ThirdImg,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}

			return nil, err
		}

		err = repo.Cache.SetStruct(add.TableName()+"_"+strconv.Itoa(id), add)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), add)
		if err != nil {
			return nil, err
		}
	}
	return add, nil
}

func (repo *AddRepository) FindAll() ([]*model.Add, error) {

	val, err := repo.Cache.Get(model.Add{}.TableName() + "_all")
	if err != nil {
		return nil, err
	}

	res := []*model.Add{}

	if val == nil || *val == "" {
		rows, err := repo.Store.Db.Query(
			fmt.Sprintf(
				"select * from %s",
				model.Add{}.TableName(),
			),
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			add := &model.Add{}
			err := rows.Scan(&add.ID,
				&add.Title,
				&add.FirstItem,
				&add.SecondItem,
				&add.ThirdItem,
				&add.FourthItem,
				&add.FirstImg,
				&add.SecondImg,
				&add.ThirdImg,
				&add.CreateTime,
				&add.UpdateTime,
			)
			if err != nil {
				return nil, err
			}
			res = append(res, add)
		}

		err = repo.Cache.SetStruct(model.Add{}.TableName()+"_all", res)
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

func (repo *AddRepository) Update(add *model.Add) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf(
			"update %s set title = $1, first_item = $2, second_item = $3, third_item = $4, fourth_item = $5, first_img = $6, second_img = $7, third_img = $8 where id = $9",
			add.TableName(),
		),
		add.Title,
		add.FirstItem,
		add.SecondItem,
		add.ThirdItem,
		add.FourthItem,
		add.FirstImg,
		add.SecondImg,
		add.ThirdImg,
		add.ID,
	)
	repo.Cache.Client.Del(add.TableName() + "_" + strconv.Itoa(add.ID))
	err = repo.Cache.Del(add.TableName() + "_all")
	if err != nil {
		return err
	}

	return nil
}

func (repo *AddRepository) Delete(id int) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf("delete from %s where id = $1", model.Add{}.TableName()),
		id,
	)
	repo.Cache.Client.Del(model.Add{}.TableName() + "_" + strconv.Itoa(id))
	err = repo.Cache.Del(model.Add{}.TableName() + "_all")
	if err != nil {
		return err
	}

	return nil
}
