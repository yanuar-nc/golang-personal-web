package repository

import "learning/personal/internal/entity"

type RepositoryInterace interface {
	ExperienceRepository
	UserRepository
}

type ExperienceRepository interface {
	FetchAll() ([]entity.Experience, error)
}

type UserRepository interface {
	FetchAll() ([]entity.Experience, error)
}
