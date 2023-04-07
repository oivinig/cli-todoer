package usecases

import "github.com/oivinig/cli-todoer/internal/app/domain"

type TaskRepository interface {
	Create(string) error
	FindAll() (domain.TodoList, error)
	DeleteById(string) error
}
