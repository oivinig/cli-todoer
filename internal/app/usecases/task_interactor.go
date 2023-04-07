package usecases

import "github.com/oivinig/cli-todoer/internal/app/domain"

// A TaskInteractor belong to the usecases layer
type TaskInteractor struct {
	TaskRepository TaskRepository
}

func (ti *TaskInteractor) List() (domain.TodoList, error) {
	tasks, err := ti.TaskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//TODO: implement next methods