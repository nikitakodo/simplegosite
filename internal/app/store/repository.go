package store

import "simplesite/internal/app/model"

type RepositoryInterface interface {
	Create(abstractModel *model.Interface) error
	Find(int) (*model.Interface, error)
	Update(abstractModel *model.Interface) error
	Delete(int) error
}

type AdminRepository interface {
	RepositoryInterface
	FindByEmail(string) (*model.Admin, error)
}
