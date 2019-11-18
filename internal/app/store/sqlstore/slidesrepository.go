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

type SlidesRepository struct {
	Store *Store
	Cache *services.Cache
}

func (repo *SlidesRepository) Create(slide *model.Slide) error {

	if err := slide.Validate(); err != nil {
		return err
	}

	err := repo.Store.Db.QueryRow(
		fmt.Sprintf(
			"INSERT INTO %s (order, first_title, second_title, third_title, img) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			slide.TableName(),
		),
		slide.Order,
		slide.FirstTitle,
		slide.SecondTitle,
		slide.ThirdTitle,
		slide.Img,
	).Scan(slide.ID)

	if err != nil {
		return err
	}

	err = repo.Cache.Del(slide.TableName() + "_all")
	if err != nil {
		return err
	}

	return nil
}

func (repo *SlidesRepository) Find(id int) (*model.Slide, error) {
	slide := &model.Slide{}

	val, err := repo.Cache.Get(slide.TableName() + "_" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	if val == nil || *val == "" {
		if err := repo.Store.Db.QueryRow(
			fmt.Sprintf(
				"SELECT id, order, first_title, second_title, third_title, img FROM %s WHERE id = $1",
				slide.TableName(),
			),
			id,
		).Scan(
			slide.ID,
			slide.Order,
			slide.FirstTitle,
			slide.SecondTitle,
			slide.ThirdTitle,
			slide.Img,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}

			return nil, err
		}

		err = repo.Cache.SetStruct(slide.TableName()+"_"+strconv.Itoa(id), slide)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), slide)
		if err != nil {
			return nil, err
		}
	}
	return slide, nil
}

func (repo *SlidesRepository) FindAll() ([]*model.Slide, error) {

	val, err := repo.Cache.Get(model.Slide{}.TableName() + "_all")
	if err != nil {
		return nil, err
	}

	res := []*model.Slide{}

	if val == nil || *val == "" {
		rows, err := repo.Store.Db.Query(
			fmt.Sprintf(
				"select * from %s order by \"order\"",
				model.Slide{}.TableName(),
			),
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			slide := &model.Slide{}
			err := rows.Scan(&slide.ID,
				&slide.FirstTitle,
				&slide.SecondTitle,
				&slide.ThirdTitle,
				&slide.Img,
				&slide.CreateTime,
				&slide.UpdateTime,
				&slide.Order,
			)
			if err != nil {
				return nil, err
			}
			res = append(res, slide)
		}

		err = repo.Cache.SetStruct(model.Slide{}.TableName()+"_all", res)
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

func (repo *SlidesRepository) Update(slide *model.Slide) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf(
			"update %s set order = $1, first_title=$2, second_title=$3, third_title=$4, img=$5 where id = $6",
			slide.TableName(),
		),
		slide.Order,
		slide.FirstTitle,
		slide.SecondTitle,
		slide.ThirdTitle,
		slide.Img,
		slide.ID,
	)
	repo.Cache.Client.Del(slide.TableName() + "_" + strconv.Itoa(slide.ID))
	err = repo.Cache.Del(slide.TableName() + "_all")
	if err != nil {
		return err
	}

	return nil
}

func (repo *SlidesRepository) Delete(id int) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf("delete from %s where id = $1", model.Slide{}.TableName()),
		id,
	)
	if err != nil {
		repo.Cache.Client.Del(model.Slide{}.TableName() + "_" + strconv.Itoa(id))
	}

	return err
}
