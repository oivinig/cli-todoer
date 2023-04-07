package domain

const (
	DONE_TASK = "DONE"
	TODO_TASK = "TODO"
)

type Task struct {
	Id string
	Description string
	// Status string //TODO_TASK OR DONE_TASK //TODO: create logic to be the name of the bucket
}

type TodoList []Task //A TodoList is a list of Tasks
