package interfaces

import (
	"github.com/oivinig/cli-todoer/internal/app/domain"
)

type TaskRepository struct {
	DBHandler DBHandler
}

// TODO: interfaces that must be met
// type TaskRepository interface {
// 	Create(domain.Task) error
// 	FindAll() (domain.TodoList, error)
// 	DeleteById(string) error //It will receive the uuid
// 	Edit(string) (string, error) 
// }

func (tr *TaskRepository) Create(t domain.Task) error {
	key := []byte(t.Id)
	value := []byte(t.Description)
	err := tr.DBHandler.Create(key, value)
	if err != nil {
		return err
	}
	return nil
}

// FIXME: implement
func (tr *TaskRepository) FindAll() (domain.TodoList, error) {
	// values := tr.DBHandler.Retrieve()
	return domain.TodoList{}, nil
}

//TODO: implement next methods