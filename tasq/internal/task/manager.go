package task

import (
	"tasq/internal/entity"
	"tasq/internal/storage"
)

type TaskManager interface {
	AddTask(t entity.Task) error
	ListTasks() ([]entity.Task, error)
}

type taskManagerImpl struct {
	storage storage.Storage
}

func NewTaskManager(s storage.Storage) TaskManager {
	return &taskManagerImpl{storage: s}
}

func (tm taskManagerImpl) AddTask(t entity.Task) error {
	return tm.storage.Save(t)
}

func (tm taskManagerImpl) ListTasks() ([]entity.Task, error) {
	return tm.storage.Load()
}
