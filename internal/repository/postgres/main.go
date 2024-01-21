package postgres

import "github.com/fixelti/family-hub/internal/repository/postgres/user"

type RepositoryManager interface{
	user.UserRepository
}

type repository struct{
	user.UserRepository
}

func New() RepositoryManager {
	return repository{}
}