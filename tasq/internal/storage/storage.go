package storage

import "tasq/internal/entity"

type Storage interface {
	Save(t entity.Task) error
	Load() ([]entity.Task, error)
	Update(t entity.Task) error
}
