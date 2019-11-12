package store

type Store interface {
	GetRepository(repositoryName string) (RepositoryInterface, error)
}
