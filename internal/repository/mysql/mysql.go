package mysql

import "learning/personal/internal/repository"

type Repository struct {
}

var _ repository.RepositoryInterace = &Repository{}

// func (r )
