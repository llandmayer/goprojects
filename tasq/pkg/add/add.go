package add

import (
	"encoding/csv"
	"errors"
	"io"
)

type Item struct {
	Name        string
	Description string
	Status      string
}

type TaskService struct {
	writer io.Writer
}

func NewTaskService(w io.Writer) *TaskService {
	return &TaskService{writer: w}
}

func (s *TaskService) Add(task Item) error {
	if task.Name == "" || task.Description == "" || task.Status == "" {
		return errors.New("task name, description and status must not be empty")
	}
	record := []string{task.Name, task.Description, task.Status}
	csvWriter := csv.NewWriter(s.writer)
	if err := csvWriter.Write(record); err != nil {
		return err
	}
	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		return err
	}
	return nil
}
