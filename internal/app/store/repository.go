package store

import "simplesite/internal/app/model"

type RepositoryInterface interface {
	Create(abstractModel *model.Interface) error
	Find(id int) (*model.Interface, error)
	FindAll() ([]*model.Interface, error)
	Update(abstractModel *model.Interface) error
	Delete(id int) error
}

type AdminRepository interface {
	RepositoryInterface
	FindByEmail(string) (*model.Admin, error)
}
