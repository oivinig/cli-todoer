package usecases

import (
	"github.com/oivinig/cli-todoer/internal/app/domain"
)

// A TaskInteractor belong to the usecases layer
type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (ti *TaskInteractor) Add(description string) error {
	err := ti.TaskRepository.Create(description)
	if err != nil {
		return err
	}
	return nil
}

func (ti *TaskInteractor) List() (domain.TodoList, error) {
	tasks, err := ti.TaskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ti *TaskInteractor) Do(key string) error {
	err := ti.TaskRepository.DeleteById(key)
	if err != nil {
		return err
	}
	return nil
}