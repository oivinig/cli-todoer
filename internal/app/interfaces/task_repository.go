package interfaces

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/oivinig/cli-todoer/internal/app/domain"
)

type TaskRepository struct {
	DBHandler DBHandler
}

func (tr *TaskRepository) Create(description string) error {
	err := tr.DBHandler.Create(description)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) DeleteById(key string) error {
	id, _ := strconv.Atoi(key)
	err := tr.DBHandler.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) FindAll() (domain.TodoList, error) {
	var todoList domain.TodoList 
	k, v, err := tr.DBHandler.Retrieve()
	if err != nil {
		return nil, err
	}
	todoList = makeTodoList(k, v)
	return todoList, nil
}

// FIXME: this should be at another layer
func makeTodoList(keys [][]byte, values [][]byte) domain.TodoList {
	var todoList []string
	for i := 0; i < len(keys); i++ {
		key := btoi(keys[i])
		value := string(values[i])
		str := fmt.Sprintf("%d. %s", key, value)
		todoList = append(todoList, str)
	}
	return todoList
}


func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}