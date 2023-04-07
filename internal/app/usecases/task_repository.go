package usecases

import "github.com/oivinig/cli-todoer/internal/app/domain"

type TaskRepository interface {
	Create(domain.Task) error
	FindAll() (domain.TodoList, error)
	DeleteById(string) error //It will receive the uuid
	Edit(string) (string, error) 
}
