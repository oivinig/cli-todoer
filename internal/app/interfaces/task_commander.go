package interfaces

import "github.com/oivinig/cli-todoer/internal/app/usecases"

type TaskCommander struct {
	TaskInteractor usecases.TaskInteractor
}

func NewTaskCommander(dbHandler DBHandler) *TaskCommander {
	return &TaskCommander{
		TaskInteractor: usecases.TaskInteractor{
			TaskRepository: &TaskRepository{
				DBHandler: dbHandler,
			},
		},
	}
}
